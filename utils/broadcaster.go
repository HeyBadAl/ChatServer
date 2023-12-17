// Package utils provide utility functions and data structures for the ChatServer Application.
package utils

type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
}

var Subscribers = make(map[chan<- Message]struct{})

// AddSubscriber add a new message channel to the list of subscribers
func AddSubscriber(subscriber chan<- Message) {
	SubscribersMutex.Lock()
	defer SubscribersMutex.Unlock()
	Subscribers[subscriber] = struct{}{}
}

// BroadcastMessage sends a message to all subscribers
func BroadcastMessage(msg Message) {
	SubscribersMutex.Lock()
	defer SubscribersMutex.Unlock()
	for ch := range Subscribers {
		ch <- msg
	}
}
