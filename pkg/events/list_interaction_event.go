package events

type ListInteractionEvent struct {
	BaseMessageEvent
}

func NewListInteractionEvent(baseMessageEvent BaseMessageEvent, text string) *ListInteractionEvent {
	return &ListInteractionEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
