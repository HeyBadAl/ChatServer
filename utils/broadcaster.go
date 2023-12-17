package utils

type Message struct {
  User string `json:"user"`
  Content string `json:"content"`
}

var Subscribers = make(map[chan<- Message]struct{})

func AddSubscriber(subscriber chan<- Message) {
	SubscribersMutex.Lock()
	defer SubscribersMutex.Unlock()
	Subscribers[subscriber] = struct{}{}
}

func BroadcastMessage(msg Message) {
	SubscribersMutex.Lock()
	defer SubscribersMutex.Unlock()
	for ch := range Subscribers {
		ch <- msg
	}
}
