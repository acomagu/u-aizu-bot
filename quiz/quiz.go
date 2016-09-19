package quiz

import (
	"github.com/acomagu/u-aizu-bot/types"
)

// Talk method start quiz game with user if sent message means "Quiz".
func Talk(chatroom types.Chatroom) {
	text := <-chatroom.In
	if text != "クイズ" {
		return
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
}

func isCorrectAnswer(text types.Message, qa QA) bool {
	return text == types.Message(qa.answer)
}
