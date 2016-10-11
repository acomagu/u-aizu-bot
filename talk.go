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
		for i, topicChan := range topicChans {
			changeDestTopicTo <- topicChan.Chatroom
			if(i > 0) {  // for the start time.
				broadcastPool <- true
			}
			didTalk := <-topicChan.Return
			if didTalk {
				clearPool <- true
				break
			}
		}
		clearPool <- true
	}
}

// This pipe stores messages from user with flowing next Chatroom(middleChatroom). And this provides functions, clearPool and broadcastPool. This is used in controller().
func poolMessages(chatroom types.Chatroom) (types.Chatroom, chan bool, chan bool) {
	middleChatroom := types.Chatroom{
		In:  make(chan types.Message),
		Out: chatroom.Out,
		Token: chatroom.Token,
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

// distributeMessage pass message from chatroom to chatroom. The chatroom of destination will change as needed, changed by value of channel, changeDestTopicTo.
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

// loopTopic just loops topic.
func loopTopic(topic types.Topic, chatroom types.Chatroom) TopicChan {
	topicChan := TopicChan{
		Chatroom: types.Chatroom{
			In:  make(chan types.Message),
			Out: chatroom.Out,
			Token: chatroom.Token,
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
