package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

func talk(chatroom types.Chatroom) {
	for {
		for _, topic := range topics {
			topic(chatroom)
		}
	}
}
