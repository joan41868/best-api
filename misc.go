package main

import (
	"fmt"
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
