package main

import (
	"github.com/acomagu/u-aizu-bot/types"
	"fmt"
)

func talk(chatroom types.Chatroom) {
	for _, topic := range topics {
		topicChatroom := types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan types.Message),
		}
		go loopTopic(topic, topicChatroom)
		go sendMessageFromTopicChatroom(chatroom, topicChatroom)
		go chainMessageFromChatroomToTopicChatroom(chatroom, topicChatroom)
	}
}

func loopTopic(topic func(types.Chatroom), topicChatroom types.Chatroom) {
	for {
		topic(topicChatroom)
	}
}

func chainMessageFromChatroomToTopicChatroom(chatroom types.Chatroom, topicChatroom types.Chatroom) {
	for {
		// topicChatroom.In <- <-chatroom.In
		tmp := <-chatroom.In
		topicChatroom.In <- tmp
		fmt.Println(tmp)
	}
}

func sendMessageFromTopicChatroom(chatroom types.Chatroom, topicChatroom types.Chatroom) {
	for {
		// chatroom.Out <- <-topicChatroom.Out
		tmp := <-topicChatroom.Out
		chatroom.Out <- tmp
		fmt.Println(tmp)
	}
}
