package quiz

import (
	"github.com/acomagu/u-aizu-bot/types"
)

// Talk method start quiz game with user if sent message means "Quiz".
func Talk(chatroom chan types.Message) {
	text := <-chatroom
	if text != "クイズ" {
		return
	}
	qa := oneQA()
	for _, message := range qa.question {
		chatroom <- message
	}

	text = <-chatroom
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
