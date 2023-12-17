package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

var (
	messages    []Message
	mutex       sync.Mutex
	upgrader    = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	subscribers = make(map[*websocket.Conn]struct{})
)

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/send", sendMessageHandler)

	go notifySubscribers()

	http.ListenAndServe(":8080", nil)
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	subscribers[conn] = struct{}{}
	mutex.Unlock()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		broadcastMessage(msg)
	}
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mutex.Lock()
	messages = append(messages, msg)
	mutex.Unlock()

	broadcastMessage(msg)

	w.WriteHeader(http.StatusNoContent)
}

func broadcastMessage(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()
	for conn := range subscribers {
		err := conn.WriteJSON(msg)
		if err != nil {
			delete(subscribers, conn)
			conn.Close()
		}
	}
}

func notifySubscribers() {
	for {
		msg := <-messages
		broadcastMessage(msg)
	}
}

