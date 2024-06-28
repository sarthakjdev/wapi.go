package events

type MessageTemplateStatusUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewMessageTemplateStatusUpdateEvent() *MessageTemplateStatusUpdateEvent {
	return &MessageTemplateStatusUpdateEvent{}
}
