package main

import (
	"fmt"

	"github.com/acomagu/u-aizu-bot/types"
)

// TopicConnection includes chatroom channel, the channel pass the returned value from topic.
type TopicConnection struct {
	Chatroom types.Chatroom
	Return   chan bool
}

func talk(chatroom types.Chatroom) {
	topicConnections := []TopicConnection{}
	for _, topic := range topics {
		ch := make(chan bool)
		topicChatroom := types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan types.Message),
		}
		go loopTopic(topic, topicChatroom, ch)
		topicConnections = append(topicConnections, TopicConnection{
			Chatroom: topicChatroom,
			Return:   ch,
		})
	}
	tryTopic := make(chan types.Chatroom)
	middleChatroom := types.Chatroom{
		In: make(chan Message),
		Out: make(chan Message),
	}
	go controller(chatroom, topicConnections, tryTopic)
	go passMessage(middleChatroom, tryTopic)
}

func controller(chatroom types.Chatroom, topicConnections []TopicConnection, tryTopic chan types.Chatroom) {
	for {
		for _, topicConnection := range topicConnections {
			tryTopic <- topicConnection.Chatroom
			didTalk := <-topicConnection.Return
			if didTalk {
				break
			}
		}
	}
}

func passMessage(middleChatroom types.Chatroom, tryTopic <-chan types.Chatroom) {
	var dest types.Chatroom
	for {
		select {
		case message := <-chatroom.In:
			if dest == (types.Chatroom{}) {
				fmt.Println("Error: the destination chatroom in not set.")
				break
			}
			dest.In <- message
		case _dest := <-tryTopic:
			dest = _dest
		}
	}
}

func loopTopic(topic types.Topic, topicChatroom types.Chatroom, didTalk chan<- bool) {
	for {
		didTalk <- topic(topicChatroom)
	}
}
