package main

import (
	"github.com/acomagu/u-aizu-bot/types"
)

func talk(chatroom chan types.Message) {
	for {
		for _, topic := range topics {
			topic(chatroom)
		}
	}
}
