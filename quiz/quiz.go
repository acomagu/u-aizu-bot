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
	chatroom.Out <- qa.question
	// for _, message := range qa.question {
	// 	chatroom.Out <- message
	// }

	text = <-chatroom.In
	var messages []types.Message

	if isCorrectAnswer(text, qa) {
		messages = append(messages,"なんで知ってるの...?")
		chatroom.Out <- messages
	} else {
		messages = append(messages,"やーいやーーいwwwwwwwwwwwwwwwwwww")
		messages = append(messages,types.Message("せぃかぃゎ" + qa.answer))
		chatroom.Out <- messages
	}
}

func isCorrectAnswer(text types.Message, qa QA) bool {
	return text == types.Message(qa.answer)
}
