package events

// MessageSentEvent represents an event indicating that a message has been sent.
type MessageSentEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
}

// NewMessageSentEvent creates a new instance of MessageSentEvent.
func NewMessageSentEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo string) *MessageSentEvent {
	return &MessageSentEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
	}
}
