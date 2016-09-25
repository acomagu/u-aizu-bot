package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

func talk(chatroom types.Chatroom) {
	var topicChatrooms = []types.Chatroom{}
	for _, topic := range topics {
		topicChatroom := types.Chatroom{
			In:  make(chan types.Message),
			Out: chatroom.Out,
		}
		go loopTopic(topic, topicChatroom)
		topicChatrooms = append(topicChatrooms, topicChatroom)
	}
	go chainMessageFromChatroomToTopicChatroom(chatroom, topicChatrooms)
}

func loopTopic(topic func(types.Chatroom), topicChatroom types.Chatroom) {
	for {
		topic(topicChatroom)
	}
}

func chainMessageFromChatroomToTopicChatroom(chatroom types.Chatroom, topicChatrooms []types.Chatroom) {
	for {
		text := <-chatroom.In
		for _, topicChatroom := range topicChatrooms {
			topicChatroom.In <- text
		}
	}
}
