package events

type LocationMessageEvent struct {
	BaseMessageEvent
}

func NewLocationMessageEvent(baseMessageEvent BaseMessageEvent, text string) *LocationMessageEvent {
	return &LocationMessageEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
