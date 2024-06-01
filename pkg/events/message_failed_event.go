package events

type MessageFailedEvent struct {
	BaseSystemEvent `json:",inline"`
	MessageId       string `json:"messageId"`
	SentTo          string `json:"sentTo"`
	FailReason      string `json:"failReason"`
}

func NewMessageFailedEvent(baseSystemEvent BaseSystemEvent, messageId, sendTo, failReason string) *MessageFailedEvent {
	return &MessageFailedEvent{
		BaseSystemEvent: baseSystemEvent,
		MessageId:       messageId,
		SentTo:          sendTo,
		FailReason:      failReason,
	}

}
