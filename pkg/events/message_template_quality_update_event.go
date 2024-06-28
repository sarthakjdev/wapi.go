package events

type MessageTemplateQualityUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewMessageTemplateQualityUpdateEvent() *MessageTemplateQualityUpdateEvent {
	return &MessageTemplateQualityUpdateEvent{}
}
