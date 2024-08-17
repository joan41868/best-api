package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Prophet interface {
	SubscribeToGod(g God)
	OnMessage(msg string)
	GetGodlyWords() chan string
}

type God interface {
	EmitWordOfGod(msg string) string
}

type Stefko struct {
	godlyWords chan string
}

func (s Stefko) OnMessage(msg string) {
	fmt.Printf("Stefko received message from God: %s\n", msg)
}

func (s Stefko) SubscribeToGod(g God) {
	go func() {
		for {
			select {
			case msg := <-s.godlyWords:
				s.OnMessage(msg)
			}
		}
	}()
}

type StefkoMuza struct {
	prophets []Prophet
}

func (s StefkoMuza) EmitWordOfGod(msg string) string {
	for _, pr := range s.prophets {
		pr.GetGodlyWords() <- msg
	}
	return msg
}

func (s Stefko) GetGodlyWords() chan string {
	return s.godlyWords
}

func NewStefko() Prophet {
	return Stefko{
		godlyWords: make(chan string),
	}
}

func NewMuza(prophets []Prophet) God {
	return StefkoMuza{
		prophets: prophets,
	}
}

func selectRandomMessage(messages []string) string {
	min := 0
	max := len(messages)
	num := rand.Intn(max-min) + min
	return messages[num]
}

var messages = []string{}

func readMessages() {
	// open data.txt and read each line, where each line is a message
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	stefko := NewStefko()
	stefkoMuza := NewMuza([]Prophet{stefko})

	stefko.SubscribeToGod(stefkoMuza)
	go func() {
		for {
			stefkoMuza.EmitWordOfGod(selectRandomMessage(messages))
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Path[1:]
		messages = append(messages, data)
	})
	http.ListenAndServe(":8080", nil)
}
