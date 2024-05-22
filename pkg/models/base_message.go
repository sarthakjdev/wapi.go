package models

type ApiCompatibleJsonConverterConfigs struct {
	ReplyToMessageId  string
	SendToPhoneNumber string
}

type Context struct {
	MessageId string `json:"message_id,omitempty"`
}

// Base API Payload to send messages
type BaseMessagePayload struct {
	Context          *Context `json:"context,omitempty"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
}

func NewBaseMessagePayload(to, messageType string) BaseMessagePayload {
	return BaseMessagePayload{
		To:               to,
		Type:             messageType,
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
	}
}

type MediaMessagePayload struct {
}

type AudioMessagePayload struct {
}

type TextMessageApiPayload struct {
	BaseMessagePayload `json:",inline"`
	Text               TextMessageApiPayloadText `json:"text" validate:"required"`
}

type BaseMessage interface {
	ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error)
}
