package repeater

import (
	"github.com/acomagu/u-aizu-bot/types"
)

// Talk is main func, just repeat.
func Talk(chatroom types.Chatroom) bool {
	chatroom.Out <- []types.Message{<-chatroom.In}
	return true
}
