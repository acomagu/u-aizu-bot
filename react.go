package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

var chatrooms = make(map[types.UserID]chan types.Message)

func react(text types.Message, userID types.UserID) error {
	chatroom, ok := chatrooms[userID]
	if !ok {
		chatroom = make(chan types.Message)
		go sendMessageFromChatroom(chatroom, userID)
		go talk(chatroom)
		chatrooms[userID] = chatroom
	}
	chatroom <- text
	return nil
}

func sendMessageFromChatroom(chatroom chan types.Message, userID types.UserID) {
	for {
		text := <-chatroom
		bot.SendText([]string{string(userID)}, string(text))
	}
}
