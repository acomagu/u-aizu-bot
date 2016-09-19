package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

func talk(chatroom types.Chatroom) {
	for _, topic := range topics {
		topicChatroom := types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan types.Message),
		}
		go topic(topicChatroom)
		go sendMessageFromTopicChatroom(chatroom, topicChatroom)
		go chainMessageFromChatroomToTopicChatroom(chatroom, topicChatroom)
	}
}

func chainMessageFromChatroomToTopicChatroom(chatroom types.Chatroom, topicChatroom types.Chatroom) {
	for {
		topicChatroom.In <- <-chatroom.In
	}
}

func sendMessageFromTopicChatroom(chatroom types.Chatroom, topicChatroom types.Chatroom) {
	for {
		chatroom.Out <- <-topicChatroom.Out
	}
}
