package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/acomagu/u-aizu-bot/chatroom"
)

type UserID string

var bot *linebot.Client

func main() {
	var err error
	configProxy()

	bot, err = lineClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	listen()
}

func listen() {
	port := os.Getenv("PORT")
	http.HandleFunc("/callback", handleRequest)
	http.ListenAndServe(":"+port, nil)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	events, err := bot.ParseRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, event := range events {
		text := ""
		if event != nil && event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				text = message.Text
			}
		}
		err = react(ReplyToken(event.ReplyToken), chatrooms.Message(text), UserID(event.Source.UserID))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func configProxy() {
	fixieURL := os.Getenv("FIXIE_URL")
	os.Setenv("HTTP_PROXY", fixieURL)
	os.Setenv("HTTPS_PROXY", fixieURL)
}

func lineClient() (*linebot.Client, error) {
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineChannelAccessToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	bot, err := linebot.New(lineChannelSecret, lineChannelAccessToken)

	return bot, err
}
