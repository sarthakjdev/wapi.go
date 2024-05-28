package events

type TextMessageEvent struct {
	BaseMessageEvent
	Text string `json:"text"`
}

func NewTextMessageEvent(messageId, from, text string) *TextMessageEvent {
	return &TextMessageEvent{
		BaseMessageEvent: BaseMessageEvent{
			MessageId: messageId,
			Context: MessageContext{
				From: from,
			},
		},
		Text: text,
	}
}
