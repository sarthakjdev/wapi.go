package manager

import "github.com/sarthakjdev/wapi.go/pkg/client"

type MessageManager struct {
	BaseManager
	client client.Client
}

func NewMessageManager() *MessageManager {
	return &MessageManager{}
}

func (mm *MessageManager) Send() {
	
	// Send message
}

func (mm *MessageManager) Reply() {
	// Reply to message
}
