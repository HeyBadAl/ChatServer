package main

import (
	"net/http"

	"github.com/gorilla/websocket"

	"HeyBadAl/ChatServer/handlers"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/send", handlers.SendMessageHandler)

	go handlers.NotifySubscribers()

	http.ListenAndServe(":8080", nil)
}
