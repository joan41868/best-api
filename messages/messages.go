package messages

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

var Messages = []string{}

func ReadMessages() {
	// open data.txt and read each line, where each line is a message
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			continue
		}
		Messages = append(Messages, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func WriteNewMessage(msg string) {
	f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString("\n" + msg + "\n"); err != nil {
		log.Fatal(err)
	}
}

func SelectRandomMessage() string {
	min := 0
	max := len(Messages)
	num := rand.Intn(max-min) + min
	return Messages[num]
}

func init() {
	ReadMessages()
}
