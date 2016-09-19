package main

import (
	"github.com/acomagu/u-aizu-bot/quiz"
	"github.com/acomagu/u-aizu-bot/timetable"
	"github.com/acomagu/u-aizu-bot/types"
)

var topics = []func(types.Chatroom){
	quiz.Talk,
	timetable.Timetable,
}
