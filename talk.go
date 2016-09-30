package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

func talk(chatroom types.Chatroom) {
	for _, topic := range topics {
		isTalked := tryTopic(topic)
		if isTalked {
			break
		}
	}
}

func tryTopic(topic func(types.Chatroom)) bool {
	topicChatroom := types.Chatroom{
		In:  make(chan types.Message),
		Out: chatroom.Out,
	}
	isTalked := make(chan bool)

	go func(topic func(types.Chatroom), topicChatroom types.Chatroom, isTalked chan bool) {
		isTalked <- topic(topicChatroom)
	}(topic, topicChatroom, isTalked)
	go passMessage(topicChatroom)

	return <-isTalked
}

func passMessage(chatroom types.Chatroom, topicChatrooms []types.Chatroom) {
	for {
		text := <-chatroom.In
		for _, topicChatroom := range topicChatrooms {
			topicChatroom.In <- text
		}
	}
}
