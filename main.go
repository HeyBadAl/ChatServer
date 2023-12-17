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
	mutext      sync.Mutex
	subscribers = make(map[chan<- Message]struct{})
	upgrader    = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
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

	// new channel for this client
	messageChan := make(chan Message)
	mutex.Lock()
	subscribers[messageChan] = struct{}{}
	mutex.Unlock()

	// Listen for messages from the client
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		messageChan <- msg
	}
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add the message to the in-memory data structure
	mutex.lock()
	messages = append(messages, msg)
	mutex.Unlock()

	// Notify all subscribers about the new message
	broadcastMessage(msg)

	w.WriteHeader(http.StatusNoContent)
}

func broadcastMessage(msg Message) {
	mutex.Lock()
	defer mutex.Unlock()
	for ch := range subscribers {
		ch <- msg
	}
}

func notifySubscribers() {
	for {
		mutex.Lock()
		if len(messages) > 0 {
			msg := messages[0]
			messages = messages[1:]
			mutex.Unlock()

			broadcastMessage(msg)
		} else {
			mutex.Unlock()
		}
	}
}
