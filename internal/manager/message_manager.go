package manager

type MessageManager struct {
	requester requestClient
}

func NewMessageManager(requester requestClient) *MessageManager {
	return &MessageManager{
		requester: requester,
	}
}

func (mm *MessageManager) Send() {
	// Send message
}

func (mm *MessageManager) Reply() {
	// Reply to message
}
