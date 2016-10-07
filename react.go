package main

import (
	"strings"

	"github.com/acomagu/u-aizu-bot/types"
)

var chatrooms = make(map[types.UserID]types.Chatroom)

// react is runneds synchronously
func react(text types.Message, userID types.UserID) error {
	chatroom, ok := chatrooms[userID]
	if !ok {
		chatroom = types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan types.Message),
		}
		go sendMessageFromChatroom(chatroom.Out, userID)
		go talk(chatroom)
		chatrooms[userID] = chatroom
	}
	chatroom.In <- text
	return nil
}

func sendMessageFromChatroom(chatroom <-chan types.Message, userID types.UserID) {
	for {
		text := <-chatroom
		bot.SendText([]string{string(userID)}, string(text))
	}
}
