package main

import (
	"fmt"
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
	logMessage("->", text)
	return nil
}

func logMessage(prefix string, text types.Message) {
	for i, line := range strings.Split(string(text), "\n") {
		if i == 0 {
			fmt.Print(prefix + " ")
		} else {
			fmt.Print(".. ")
		}
		fmt.Println(line)
	}
}

func sendMessageFromChatroom(chatroom <-chan types.Message, userID types.UserID) {
	for {
		text := <-chatroom
		bot.SendText([]string{string(userID)}, string(text))

		logMessage("<-", text)
	}
}
