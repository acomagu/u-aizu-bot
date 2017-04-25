package repeater

import (
	"github.com/acomagu/u-aizu-bot/chatrooms"
)

// Talk is main func, just repeat.
func Talk(chatroom chatrooms.Room) bool {
	chatroom.Out <- []chatrooms.Message{<-chatroom.In}
	return true
}
