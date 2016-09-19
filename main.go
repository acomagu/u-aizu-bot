package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"net/http"
	"os"
	"strconv"
	"github.com/acomagu/u-aizu-bot/types"
)

var bot *linebot.Client

func main() {
	var err error
	configProxy()

	bot, err = lineClient()
	if err != nil {
		fmt.Println(err)
	}

	listen()
}

func listen() {
	port := os.Getenv("PORT")
	http.HandleFunc("/callback", handleRequest)
	http.ListenAndServe(":"+port, nil)
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)

	received, err := bot.ParseRequest(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, result := range received.Results {
		content := result.Content()
		if content != nil && content.IsMessage && content.ContentType == linebot.ContentTypeText {
			text, err := content.TextContent()
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(text.Text)

			err = react(types.Message(text.Text), types.UserID(content.From))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func configProxy() {
	fixieURL := os.Getenv("FIXIE_URL")
	os.Setenv("HTTP_PROXY", fixieURL)
	os.Setenv("HTTPS_PROXY", fixieURL)
}

func lineClient() (*linebot.Client, error) {
	lineChannelID, err := strconv.Atoi(os.Getenv("LINE_CHANNEL_ID"))
	if err != nil {
		return nil, err
	}
	lineChannelSecret := os.Getenv("LINE_CHANNEL_SECRET")
	lineMID := os.Getenv("LINE_MID")
	bot, err := linebot.NewClient(int64(lineChannelID), lineChannelSecret, lineMID)
	return bot, err
}
