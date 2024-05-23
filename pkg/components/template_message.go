package components

type TemplateMessage struct {
}

type TemplateMessageApiPayload struct {
	BaseMessagePayload
	Template TemplateMessage `json:"template" validate:"required"`
}

func NewTemplateMessage() (*TemplateMessage, error) {

	return &TemplateMessage{}, nil

}
