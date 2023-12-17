package handlers

import (
	"fmt"
	"net/http"

	"HeyBadAl/ChatServer/utils"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := utils.UpgradeToWebSocket(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	messageChan := make(chan utils.Message)
	utils.AddSubscriber(messageChan)

	for {
		var msg utils.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		messageChan <- msg
	}
}
