package handlers

import (
	"HeyBadAl/ChatServer/utils"
)

func NotifySubscribers() {
	for {
		utils.MessagesMutex.Lock()
		if len(utils.Messages) > 0 {
			msg := utils.Messages[0]
			utils.Messages = utils.Messages[1:]
			utils.MessagesMutex.Unlock()

			utils.BroadcastMessage(msg)
		} else {
			utils.MessagesMutex.Unlock()
		}
	}
}
