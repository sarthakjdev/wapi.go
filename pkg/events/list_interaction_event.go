package events

// ListInteractionEvent represents an interaction event related to a list.
type ListInteractionEvent struct {
	BaseMessageEvent `json:",inline"`
	Title            string `json:"title"`
	ListId           string `json:"list_id"`
	Description      string `json:"description"`
}

// NewListInteractionEvent creates a new ListInteractionEvent instance.
func NewListInteractionEvent(baseMessageEvent BaseMessageEvent, title, listId, description string) *ListInteractionEvent {
	return &ListInteractionEvent{
		BaseMessageEvent: baseMessageEvent,
		Title:            title,
		ListId:           listId,
		Description:      description,
	}
}
