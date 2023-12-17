package utils

var Messages []Message

func AddMessage(msg Message) {
	MessagesMutex.Lock()
	defer MessagesMutex.Unlock()
	Messages = append(Messages, msg)
}
