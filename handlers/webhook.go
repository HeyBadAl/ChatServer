// Package handlers provides HTTP handlers for the ChatServer application.
package handlers

import (
	"log"
	"net/http"

	"HeyBadAl/ChatServer/utils"
)

// WebhookHandler handles WebSocket connections and process incoming messages
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := utils.UpgradeToWebSocket(w, r)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	messageChan := make(chan utils.Message)
	utils.AddSubscriber(messageChan)

	for {
		var msg utils.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading WebSocket message: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		messageChan <- msg
	}
}
