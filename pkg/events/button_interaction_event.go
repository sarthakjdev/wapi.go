package events

type ButtonInteractionMessageEvent struct {
	BaseMessageEvent
}

func NewButtonInteractionEvent(baseMessageEvent BaseMessageEvent, text string) *ButtonInteractionMessageEvent {
	return &ButtonInteractionMessageEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
