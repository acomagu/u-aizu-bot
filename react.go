package main

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/acomagu/u-aizu-bot/chatrooms"
	"github.com/acomagu/u-aizu-bot/emptyroomsearching"
	"github.com/acomagu/u-aizu-bot/quiz"
	"github.com/acomagu/u-aizu-bot/timetable"
	"github.com/acomagu/u-aizu-bot/repeater"
)

// ReplyToken type is used for Reply-Token of line.(This value is essential for sending message to user by LINE.)
type ReplyToken string

type User struct {
	Cr chatrooms.Chatrooms
	ReplyTokenChan chan ReplyToken
	Name string
}

var users = make(map[UserID]User)

var replyTokenChans = make(map[UserID]chan ReplyToken)

func react(token ReplyToken, text chatrooms.Message, userID UserID) error {
	user, ok := users[userID]

	if !ok {
		user = User{
			Cr: chatrooms.New([]chatrooms.Topic{
				// user.returnName,
				emptyroomsearching.Emptyroomsearching,
				timetable.Timetable,
				quiz.Talk,
				repeater.Talk,
			}),
			ReplyTokenChan: make(chan ReplyToken),
			Name: "aaa",
		}
		users[userID] = user

		go sendMessageFromChatroom(user.ReplyTokenChan, user.Cr.Entry.Out)
	}
	user.ReplyTokenChan <- ReplyToken(token)

	user.Cr.Entry.In <- text

	logReceiving(text)

	return nil
}

func (user *User) returnName(cr chatrooms.Room) bool {
	_ = <-cr.In
	cr.Out <- []chatrooms.Message{chatrooms.Message(user.Name)}
	fmt.Printf("%+v\n", user)
	return true
}

func sendMessageFromChatroom(token <-chan ReplyToken, chatroom <-chan []chatrooms.Message) {
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

func logSending(texts []chatrooms.Message) {
	for _, text := range texts {
		logMessage("<-", text)
	}
}

func logReceiving(text chatrooms.Message) {
	logMessage("->", text)
}

func logMessage(prefix string, text chatrooms.Message) {
	for i, line := range strings.Split(string(text), "\n") {
		if i == 0 {
			fmt.Print(prefix + " ")
		} else {
			fmt.Print(".. ")
		}
		fmt.Println(line)
	}
}
