package main

import (
	"net/http"

	"HeyBadAl/ChatServer/handlers"
)


func main() {
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/send", handlers.SendMessageHandler)

	go handlers.NotifySubscribers()

	http.ListenAndServe(":8080", nil)
}
