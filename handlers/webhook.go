package handlers

import (
	"log"
	"net/http"

	"HeyBadAl/ChatServer/utils"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := utils.UpgradeToWebSocket(w, r)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
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
			return
		}
		messageChan <- msg
	}
}
