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
			In: make(chan types.Message),
			Out: make(chan types.Message),
		}
		go loopTopic(topic, topicChatroom, ch)
		topicConnections = append(topicConnections, TopicConnection{
			Chatroom: topicChatroom,
			Return:   ch,
		})
	}
	destTopicChatroom := make(chan types.Chatroom)
	go controller(chatroom, topicConnections, destTopicChatroom)
	go passMessage(chatroom, destTopicChatroom)
}

func controller(chatroom types.Chatroom, topicConnections []TopicConnection, destTopicChatroom chan types.Chatroom) {
	for {
		for _, topicConnection := range topicConnections {
			destTopicChatroom <- topicConnection.Chatroom
			didTalk := <-topicConnection.Return
			if didTalk {
				break
			}
		}
	}
}

func passMessage(chatroom types.Chatroom, destTopicChatroom <-chan types.Chatroom) {
	var dest types.Chatroom
	for {
		select {
		case message := <-chatroom.Out:
			fmt.Println("Error: the destination chatroom in not set.")
			dest.In <- message
		case _dest := <-destTopicChatroom:
			dest = _dest
		}
	}
}

func loopTopic(topic types.Topic, topicChatroom types.Chatroom, didTalk chan<- bool) {
	for {
		didTalk <- topic(topicChatroom)
	}
}
