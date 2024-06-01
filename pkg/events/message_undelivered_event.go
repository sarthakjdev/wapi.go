package events

// MessageUndeliveredEvent represents an event related to an undelivered message.
type MessageUndeliveredEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
}

// NewMessageUndeliveredEvent creates a new instance of MessageUndeliveredEvent.
func NewMessageUndeliveredEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo string) *MessageUndeliveredEvent {
	return &MessageUndeliveredEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
	}
}
