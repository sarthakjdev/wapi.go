package events

type TextMessageEvent struct {
	BaseMessageEvent
	Text string `json:"text"`
}

func NewTextMessageEvent(baseMessageEvent BaseMessageEvent, text string) *TextMessageEvent {
	return &TextMessageEvent{
		BaseMessageEvent: baseMessageEvent,
		Text:             text,
	}
}
