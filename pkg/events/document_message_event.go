package events

type DocumentMessageEvent struct {
	BaseMessageEvent
}

func NewDocumentMessageEvent(baseMessageEvent BaseMessageEvent, text string) *DocumentMessageEvent {
	return &DocumentMessageEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
