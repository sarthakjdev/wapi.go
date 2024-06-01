package events

// QuickReplyButtonInteractionEvent represents an event triggered when a user interacts with a quick reply button.
type QuickReplyButtonInteractionEvent struct {
	BaseMessageEvent `json:",inline"`
	ButtonText       string `json:"button_text"`
	ButtonPayload    string `json:"button_payload"`
}

// NewQuickReplyButtonInteractionEvent creates a new instance of QuickReplyButtonInteractionEvent.
func NewQuickReplyButtonInteractionEvent(baseMessageEvent BaseMessageEvent, buttonText, buttonPayload string) *QuickReplyButtonInteractionEvent {
	return &QuickReplyButtonInteractionEvent{
		BaseMessageEvent: baseMessageEvent,
		ButtonText:       buttonText,
		ButtonPayload:    buttonPayload,
	}
}
