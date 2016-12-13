package main

import (
	"fmt"
	"strings"
	"errors"

	"github.com/acomagu/u-aizu-bot/types"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ReplyToken type is used for Reply-Token of line.(This value is essential for sending message to user by LINE.)
type ReplyToken string

var chatrooms = make(map[types.UserID]types.Chatroom)
var replyTokenChans = make(map[types.UserID]chan ReplyToken)

// react is passed each message initialy. This is runned synchronously.
func react(token ReplyToken, text types.Message, userID types.UserID) error {
	if text == "" {
		return errors.New("react: Received empty message")
	}
	chatroom, ok1 := chatrooms[userID]
	replyTokenChan, ok2 := replyTokenChans[userID]

	// When receive message from new user
	if !ok1 || !ok2 {
		chatroom = types.Chatroom{
			In:  make(chan types.Message),
			Out: make(chan []types.Message),
		}
		replyTokenChan = make(chan ReplyToken)

		go sendMessageFromChatroom(replyTokenChan, chatroom.Out)
		go talk(chatroom)
		chatrooms[userID] = chatroom
		replyTokenChans[userID] = replyTokenChan
	}

	replyTokenChan <- ReplyToken(token)
	chatroom.In <- text

	logReceiving(text)
	return nil
}

func sendMessageFromChatroom(token <-chan ReplyToken, chatroom <-chan []types.Message) {
	for {
		// Receive this token when receive message by LINE.
		replytoken := <-token

		texts := <-chatroom
		var s []linebot.Message
		for _, text := range texts {
			s = append(s, linebot.NewTextMessage(string(text)))
		}
		if _, err := bot.ReplyMessage(string(replytoken), s...).Do(); err != nil {
			fmt.Println(err)
		}

		logSending(texts)
	}
}

func logSending(texts []types.Message) {
	for _, text := range texts {
		logMessage("<-", text)
	}
}

func logReceiving(text types.Message) {
	logMessage("->", text)
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
