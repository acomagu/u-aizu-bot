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
		chatroom = logMessagePipe(chatroom)
		go sendMessageFromChatroom(chatroom.Out, userID)
		go talk(chatroom)
		chatrooms[userID] = chatroom
	}

	chatroom.In <- text
	return nil
}

func logMessagePipe(chatroom types.Chatroom) types.Chatroom {
	out := types.Chatroom{
		In:  make(chan types.Message),
		Out: make(chan types.Message),
	}

	go func(chatroom types.Chatroom, out types.Chatroom) {
		select {
		case text := <-chatroom.In:
			logMessage("->", text)
			out.In <- text

		case text := <-chatroom.Out:
			logMessage("<-", text)
			out.Out <- text
		}
	}(chatroom, out)

	return out
}

func logMessage(prefix string, text types.Message) {
	for i, line := range strings.SplitAfter(string(text), "\n") {
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

		fmt.Println("<- " + text)
	}
}
