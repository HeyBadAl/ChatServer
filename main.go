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
	messagesMu  sync.Mutex
	subscribers = make(map[chan<- Message]struct{})
	subMu       sync.Mutex
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

	messageChan := make(chan Message)
	addSubscriber(messageChan)

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

	addMessage(msg)
	broadcastMessage(msg)

	w.WriteHeader(http.StatusNoContent)
}

func addSubscriber(subscriber chan<- Message) {
	subMu.Lock()
	defer subMu.Unlock()
	subscribers[subscriber] = struct{}{}
}

func addMessage(msg Message) {
	messagesMu.Lock()
	defer messagesMu.Unlock()
	messages = append(messages, msg)
}

func broadcastMessage(msg Message) {
	subMu.Lock()
	defer subMu.Unlock()
	for ch := range subscribers {
		ch <- msg
	}
}

func notifySubscribers() {
	for {
		messagesMu.Lock()
		if len(messages) > 0 {
			msg := messages[0]
			messages = messages[1:]
			messagesMu.Unlock()

			broadcastMessage(msg)
		} else {
			messagesMu.Unlock()
		}
	}
}
