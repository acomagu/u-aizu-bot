package quiz

import (
	"math/rand"
	"github.com/acomagu/u-aizu-bot/chatrooms"
)

// QA struct enclose the question and the answer.
type QA struct {
	question []chatrooms.Message
	answer   string
}

var qas = []QA{
	QA{
		question: []chatrooms.Message{
			"もんだぃ。ゎたしゎなんさぃ?",
			"1. 14さぃ",
			"2. 24さぃ",
			"3. 64さぃ",
		},
		answer: "3",
	},
	QA{
		question: []chatrooms.Message{
			"こんにちは。僕のラッキーカラーは何でしょう?",
			"1. Blue",
			"2. イエロー☆",
			"3. Red",
		},
		answer: "1",
	},
	QA{
		question: []chatrooms.Message{
			"やっほー",
			"1. うっほー",
			"2. ごっほー",
			"3. サッポー",
		},
		answer: "2",
	},
}

func oneQA() QA {
	return qas[rand.Intn(len(qas))]
}
