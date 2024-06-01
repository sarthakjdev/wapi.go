package components

// TemplateMessage represents a template message.
type TemplateMessage struct {
}

// TemplateMessageApiPayload represents the API payload for a template message.
type TemplateMessageApiPayload struct {
	BaseMessagePayload
	Template TemplateMessage `json:"template" validate:"required"`
}

// NewTemplateMessage creates a new instance of TemplateMessage.
func NewTemplateMessage() (*TemplateMessage, error) {
	return &TemplateMessage{}, nil
}
