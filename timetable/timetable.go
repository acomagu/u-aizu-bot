package timetable

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/acomagu/u-aizu-bot/types"
)

type namegetter struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type person struct {
	No    string    `json:"No"`
	M     [6]string `json:"M"`
	Tu    [6]string `json:"Tu"`
	W     [6]string `json:"W"`
	T     [6]string `json:"T"`
	F     [6]string `json:"F"`
	Ather string    `json:"ather"`
}

func rtClass(menber string) [6]string {
	Mon := time.Date(2016, 5, 9, 0, 0, 0, 0, time.Local)
	Tus := time.Date(2016, 5, 10, 0, 0, 0, 0, time.Local)
	Wen := time.Date(2016, 5, 11, 0, 0, 0, 0, time.Local)
	Thu := time.Date(2016, 5, 12, 0, 0, 0, 0, time.Local)
	Fre := time.Date(2016, 5, 13, 0, 0, 0, 0, time.Local)

	file, err := ioutil.ReadFile("./json/jyu2.json")
	var datasets []person
	jsonErr := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", jsonErr)
	}
	var T [6]string

	for k := range datasets {
		if datasets[k].No == menber {
			now := time.Now()
			if now.Weekday() == Mon.Weekday() {
				T = datasets[k].M
				break
			} else if Tus.Weekday() == now.Weekday() {
				T = datasets[k].Tu
				break
			} else if Wen.Weekday() == now.Weekday() {
				T = datasets[k].W
				break
			} else if Thu.Weekday() == now.Weekday() {
				T = datasets[k].T
				break
			} else if Fre.Weekday() == now.Weekday() {
				T = datasets[k].F
				break
			} else {
				break
			}
		}
	}
	T = chName(T)
	for f := range T {
		if T[f] == "" {
			T[f] = "[あき]"
		}
	}
	return T
}

func serect(menber string, dotw string) [6]string {
	// Mon := time.Date(2016, 5, 9, 0, 0, 0, 0, time.Local)
	// Tus := time.Date(2016, 5, 10, 0, 0, 0, 0, time.Local)
	// Wen := time.Date(2016, 5, 11, 0, 0, 0, 0, time.Local)
	// Thu := time.Date(2016, 5, 12, 0, 0, 0, 0, time.Local)
	// Fre := time.Date(2016, 5, 13, 0, 0, 0, 0, time.Local)

	file, err := ioutil.ReadFile("./json/jyu2.json")
	var datasets []person
	jsonerr := json.Unmarshal(file, &datasets)
	if err != nil {
		fmt.Println("Format Error: ", jsonerr)
	}
	var T [6]string

	for k := range datasets {
		if datasets[k].No == menber {
			// now := time.Now()
			if dotw == "月" {
				T = datasets[k].M
				break
			} else if dotw == "火" {
				T = datasets[k].Tu
				break
			} else if dotw == "水" {
				T = datasets[k].W
				break
			} else if dotw == "木" {
				T = datasets[k].T
				break
			} else if dotw == "金" {
				T = datasets[k].F
				break
			} else {
				break
			}
		}
	}
	T = chName(T)
	for f := range T {
		if T[f] == "" {
			T[f] = "[あき]"
		}
	}
	return T
}
func chName(code [6]string) [6]string {
	file, err := ioutil.ReadFile("./json/subjects.json")
	var datasets []namegetter
	jsonErr := json.Unmarshal(file, &datasets)
	if err != nil {
		log.Print("Format Error: ", jsonErr)
	}

	for l := range datasets {
		for i := 0; i < 6; i++ {
			if code[i] == datasets[l].Code {
				code[i] = datasets[l].Name
			}
		}
	}
	fmt.Println(code)
	return code
}

//Timetable ...
func Timetable(chatroom types.Chatroom) bool {
	text := <-chatroom.In
	ans := false
	if (text[0] == 's') || (text[0] == 'm') {
		text2 := string(text)
		words := strings.Fields(text2)
		log.Print(words[1])
		// m := rtClass(text2)
		m := serect(words[0], words[1])
		t := strings.Join(m[:], "\n")
		chatroom.Out <- types.Message(t)
		ans = true
		return ans
	}
	return ans
}
