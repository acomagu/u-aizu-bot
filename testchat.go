package main

import (
	"github.com/acomagu/u-aizu-bot/types"
	"fmt"
)

func main() {
	chatroom = types.Chatroom{}
	go sendMessageFromChatroom(chatroom.Out, userID)
	go talk(chatroom)
	chatroom.In <- text
	return nil
}

func sendMessageFromChatroom(chatroom <-chan types.Message, userID types.UserID) {
	for {
		text := <-chatroom
		fmt.Println(text)
	}
}
