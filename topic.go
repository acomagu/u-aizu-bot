package main

import (
	"github.com/acomagu/u-aizu-bot/timetable"
	"github.com/acomagu/u-aizu-bot/types"
)

var topics = []func(chan types.Message){
	quiz,
	timetable.Timetable,
}
