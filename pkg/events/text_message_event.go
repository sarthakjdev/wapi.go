package events

// TextMessageEvent represents an event for a text message.
type TextMessageEvent struct {
	BaseMessageEvent `json:",inline"`
	Text             string `json:"text"`
}

// NewTextMessageEvent creates a new TextMessageEvent instance.
func NewTextMessageEvent(baseMessageEvent BaseMessageEvent, text string) *TextMessageEvent {
	return &TextMessageEvent{
		BaseMessageEvent: baseMessageEvent,
		Text:             text,
	}
}
