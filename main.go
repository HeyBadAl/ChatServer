package main

import (
	"log"
	"net/http"

	"HeyBadAl/ChatServer/handlers"
)

func main() {
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/send", handlers.SendMessageHandler)

	go handlers.NotifySubscribers()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
