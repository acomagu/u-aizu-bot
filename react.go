package main

import (
	"fmt"
	"strings"

	"github.com/acomagu/u-aizu-bot/types"
	"github.com/line/line-bot-sdk-go/linebot"
)

var chatrooms = make(map[types.UserID]types.Chatroom)
var replyTokenChans = make(map[types.UserID]chan types.ReplyToken)

// react is runned synchronously
func react(token string, text types.Message, userID types.UserID) error {
	chatroom, ok1 := chatrooms[userID]
	replyTokenChan, ok2 := replyTokenChans[userID]

	// When receive message from new user
	if !ok1 || !ok2 {
		chatroom = types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan []types.Message),
		}
		replyTokenChan = make(chan types.ReplyToken)

		go sendMessageFromChatroom(replyTokenChan, chatroom.Out)
		go talk(chatroom)
		chatrooms[userID] = chatroom
		replyTokenChans[userID] = replyTokenChan
	}

	replyTokenChan <- types.ReplyToken(token)
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

func sendMessageFromChatroom(token <-chan types.ReplyToken, chatroom <-chan []types.Message) {
	for {
		replytoken := <-token
		text := <-chatroom
		var s []linebot.Message
		for _, content := range text {
			s = append(s, linebot.NewTextMessage(string(content)))
			logMessage("<-", content)
		}
		if _, err := bot.ReplyMessage(string(replytoken), s...).Do(); err != nil {
			fmt.Println(err)
		}
	}
}
