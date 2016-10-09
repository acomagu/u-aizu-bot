package main

import (
	"github.com/acomagu/u-aizu-bot/types"
	"github.com/line/line-bot-sdk-go/linebot"
	"fmt"
)

var chatrooms = make(map[types.UserID]types.Chatroom)

// react is runned synchronously
func react(token string,text types.Message, userID types.UserID) error {

	chatroom, ok := chatrooms[userID]
	if !ok {
		chatroom = types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan []types.Message),
		}
		go sendMessageFromChatroom(token,chatroom.Out)
		go talk(chatroom)
		chatrooms[userID] = chatroom
	}
	chatroom.In <- text
	return nil
}

func sendMessageFromChatroom(token string, chatroom <-chan []types.Message) {
	for {
		text := <-chatroom
		var s []linebot.Message
		for _, content := range text{
			s = append(s,linebot.NewTextMessage(string(content)))
		}
		// bot.SendText([]string{string(userID)}, string(text))
		if _, err := bot.ReplyMessage(token,s...).Do(); err != nil {
			fmt.Println(err)
		}
	}
}
