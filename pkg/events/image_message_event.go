package events

type ImageMessageEvent struct {
	BaseMessageEvent
}

func NewImageMessageEvent(baseMessageEvent BaseMessageEvent, text string) *ImageMessageEvent {
	return &ImageMessageEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
