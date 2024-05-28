package events

import requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"

type TextMessageEvent struct {
	BaseMessageEvent
	Text string `json:"text"`
}

func NewTextMessageEvent(messageId, from, text string, requester requestclient.RequestClient) *TextMessageEvent {
	return &TextMessageEvent{
		BaseMessageEvent: NewBaseMessageEvent(messageId, from, requester),
		Text:             text,
	}
}
