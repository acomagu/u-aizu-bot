package main

import (
	"fmt"

	"github.com/acomagu/u-aizu-bot/types"
)

// TopicChan includes chatroom channel, the channel pass the returned value from topic.
type TopicChan struct {
	Chatroom types.Chatroom
	Return   chan bool
}

func talk(chatroom types.Chatroom) {
	topicChans := []TopicChan{}
	for _, topic := range topics {
		topicChan := loopTopic(topic, chatroom)
		topicChans = append(topicChans, topicChan)
	}
	middleChatroom, clearPool, broadcastPool := poolMessages(chatroom)
	changeDestTopicTo := distributeMessage(middleChatroom)
	go controller(topicChans, changeDestTopicTo, broadcastPool, clearPool)
}

func controller(topicChans []TopicChan, changeDestTopicTo chan types.Chatroom, broadcastPool chan bool, clearPool chan bool) {
	for {
		for _, topicChan := range topicChans {
			changeDestTopicTo <- topicChan.Chatroom
			broadcastPool <- true
			didTalk := <-topicChan.Return
			if didTalk {
				clearPool <- true
				break
			}
		}
	}
}

func poolMessages(chatroom types.Chatroom) (types.Chatroom, chan bool, chan bool) {
	middleChatroom := types.Chatroom{
		In:  make(chan types.Message),
		Out: chatroom.Out,
	}
	clearPool := make(chan bool)
	broadcastPool := make(chan bool)

	go func(chatroom types.Chatroom, middleChatroom types.Chatroom, clearPool <-chan bool, broadcastPool <-chan bool) {
		pool := []types.Message{}
		for {
			select {
			case message := <-chatroom.In:
				pool = append(pool, message)
				middleChatroom.In <- message

			case <-clearPool:
				pool = []types.Message{}

			case <-broadcastPool:
				for _, message := range pool {
					middleChatroom.In <- message
				}
			}
		}
	}(chatroom, middleChatroom, clearPool, broadcastPool)

	return middleChatroom, clearPool, broadcastPool
}

func distributeMessage(middleChatroom types.Chatroom) chan types.Chatroom {
	changeDestTopicTo := make(chan types.Chatroom)

	go func(middleChatroom types.Chatroom, changeDestTopicTo <-chan types.Chatroom) {
		var dest types.Chatroom
		dest = <-changeDestTopicTo
		for {
			select {
			case message := <-middleChatroom.In:
				if dest == (types.Chatroom{}) {
					fmt.Println("Error: the destination chatroom is not set.")
					break
				}
				dest.In <- message

			case _dest := <-changeDestTopicTo:
				dest = _dest
			}
		}
	}(middleChatroom, changeDestTopicTo)

	return changeDestTopicTo
}

func loopTopic(topic types.Topic, chatroom types.Chatroom) TopicChan {
	topicChan := TopicChan{
		Chatroom: types.Chatroom{
			In:  make(chan types.Message),
			Out: chatroom.Out,
		},
		Return: make(chan bool),
	}

	go func(topic types.Topic, topicChan TopicChan) {
		for {
			topicChan.Return <- topic(topicChan.Chatroom)
		}
	}(topic, topicChan)

	return topicChan
}
