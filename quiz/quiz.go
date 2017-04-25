package quiz

import (
	"github.com/acomagu/u-aizu-bot/chatrooms"
)

// Talk method start quiz game with user if sent message means "Quiz".
func Talk(chatroom chatrooms.Room) bool {
	text := <-chatroom.In
	if text != "クイズ" {
		return false
	}
	qa := oneQA()
	chatroom.Out <- qa.question
	// for _, message := range qa.question {
	// 	chatroom.Out <- message
	// }

	text = <-chatroom.In
	var messages []chatrooms.Message

	if isCorrectAnswer(text, qa) {
		messages = append(messages,"なんで知ってるの...?")
		chatroom.Out <- messages
	} else {
		messages = append(messages,"やーいやーーいwwwwwwwwwwwwwwwwwww")
		messages = append(messages,chatrooms.Message("せぃかぃゎ" + qa.answer))
		chatroom.Out <- messages
	}
	return true
}

func isCorrectAnswer(text chatrooms.Message, qa QA) bool {
	return text == chatrooms.Message(qa.answer)
}
