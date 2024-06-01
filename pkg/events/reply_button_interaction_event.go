package events

// ReplyButtonInteractionEvent represents an interaction event triggered by a reply button.
type ReplyButtonInteractionEvent struct {
	BaseMessageEvent `json:",inline"`
	Title            string `json:"title"`
	ButtonId         string `json:"button_id"`
}

// NewReplyButtonInteractionEvent creates a new instance of ReplyButtonInteractionEvent.
func NewReplyButtonInteractionEvent(baseMessageEvent BaseMessageEvent, title, buttonId string) *ReplyButtonInteractionEvent {
	return &ReplyButtonInteractionEvent{
		BaseMessageEvent: baseMessageEvent,
		Title:            title,
		ButtonId:         buttonId,
	}
}
