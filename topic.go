package main

import (
	"github.com/acomagu/u-aizu-bot/emptyroomsearching"
	"github.com/acomagu/u-aizu-bot/quiz"
	"github.com/acomagu/u-aizu-bot/timetable"
	"github.com/acomagu/u-aizu-bot/types"
)

var topics = []types.Topic{
	quiz.Talk,
	timetable.Timetable,
	emptyroomsearching.Emptyroomsearching,
}
