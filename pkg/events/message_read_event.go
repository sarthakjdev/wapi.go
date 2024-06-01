package events

// MessageReadEvent represents an event indicating that a message has been read.
type MessageReadEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
}

// NewMessageReadEvent creates a new instance of MessageReadEvent.
func NewMessageReadEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo string) *MessageReadEvent {
	return &MessageReadEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
	}
}
