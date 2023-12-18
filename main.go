// Package main provide the entry points and configuration for the ChatServer Application.
package main

import (
	"fmt"
	"log"
	"net/http"

	"HeyBadAl/ChatServer/config"
	"HeyBadAl/ChatServer/handlers"
)

// DefaultPort is the default port for the server
const DefaultPort = "8080"

// main is the entry point for the application
func main() {
	configData := config.ReadConfig()

	http.HandleFunc("/webhook", handlers.WebhookHandler)
	http.HandleFunc("/send", handlers.SendMessageHandler)

	go handlers.NotifySubscribers()

	addr := fmt.Sprintf(":%s", configData.Port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
