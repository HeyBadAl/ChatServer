package handlers

import (
	"encoding/json"
	"net/http"

	"HeyBadAl/ChatServer/utils"
)

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg utils.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.AddMessage(msg)
	utils.BroadcastMessage(msg)

	w.WriteHeader(http.StatusNoContent)
}
