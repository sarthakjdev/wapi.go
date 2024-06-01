package events

// MessageDeliveredEvent represents an event related to an undelivered message.
type MessageDeliveredEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
}

// MessageDeliveredEvent creates a new instance of MessageUndeliveredEvent.
func NewMessageDeliveredEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo string) *MessageDeliveredEvent {
	return &MessageDeliveredEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
	}
}
