package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joan41868/best-api/messages"
)

// Upgrader to upgrade HTTP connection to WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		select {
		case <-ticker.C:
			message := messages.SelectRandomMessage()
			if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println("Error writing message:", err)
				return
			}
		}
	}
}

func main() {
	messages.ReadMessages()
	srv := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := r.URL.Path[1:]
		if data == "" || data == "favicon.ico" {
			return
		}
		messages.WriteNewMessage(data)
		messages.ReadMessages()
	})

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, messages.SelectRandomMessage())
	})

	http.HandleFunc("/ws", handleConnections)

	srv.ListenAndServe()
}
