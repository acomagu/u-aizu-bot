package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

func quiz(chatroom chan types.Message) {
	<-chatroom
	qa := oneQA()
	for _, message := range qa.question {
		chatroom <- message
	}

	text := <-chatroom
	if isCorrectAnswer(text, qa) {
		chatroom <- "なんで知ってるの...?"
	} else {
		chatroom <- "やーいやーーいwwwwwwwwwwwwwwwwwww"
		chatroom <- types.Message("せぃかぃゎ" + qa.answer)
	}
}

func isCorrectAnswer(text types.Message, qa QA) bool {
	return text == types.Message(qa.answer)
}
