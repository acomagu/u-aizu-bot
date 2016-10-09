package quiz

import (
	"github.com/acomagu/u-aizu-bot/types"
)

// Talk method start quiz game with user if sent message means "Quiz".
func Talk(chatroom types.Chatroom) bool {
	text := <-chatroom.In
	if text != "クイズ" {
		return false
	}
	qa := oneQA()
	for _, message := range qa.question {
		chatroom.Out <- message
	}

	text = <-chatroom.In
	if isCorrectAnswer(text, qa) {
		chatroom.Out <- "なんで知ってるの...?"
	} else {
		chatroom.Out <- "やーいやーーいwwwwwwwwwwwwwwwwwww"
		chatroom.Out <- types.Message("せぃかぃゎ" + qa.answer)
	}
	return true
}

func isCorrectAnswer(text types.Message, qa QA) bool {
	return text == types.Message(qa.answer)
}
