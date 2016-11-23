package verify

import (
	"regexp"

	"github.com/acomagu/u-aizu-bot/types"
)

func Talk(chatroom types.Chatroom) bool {
	received := <-chatroom.In
	if state.didVerified {
		return false
	}

	chatroom.Out <- []types.Message{
		"こんにちは",
		"よねだよっ",
		"まず、学籍番号を教えてね",
	}
	var studentID string
	for {
		received = <-chatroom.In
		studentID := regexp.MustCompile(`[sm]\d+`).FindString(string(received))
		if studentID == "" {
			break
		}
	}
	emailAddress := studentID + "@u-aizu.ac.jp"
	chatroom.Out <- []types.Message{
		types.Message("ありがとう!" + emailAddress + "に認証メールを送ったから、リンクを押してね。"),
		types.Message("確認が終わったら使えるようになるよ!"),
	}
	return true
}
