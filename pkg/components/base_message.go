package components

type MessageType string

const (
	MessageTypeLocation    MessageType = "location"
	MessageTypeAudio       MessageType = "audio"
	MessageTypeVideo       MessageType = "video"
	MessageTypeDocument    MessageType = "document"
	MessageTypeText        MessageType = "text"
	MessageTypeContact     MessageType = "contacts"
	MessageTypeInteractive MessageType = "interactive"
	MessageTypeTemplate    MessageType = "template"
	MessageTypeReaction    MessageType = "reaction"
	MessageTypeSticker     MessageType = "sticker"
	MessageTypeImage       MessageType = "image"
)

type ApiCompatibleJsonConverterConfigs struct {
	ReplyToMessageId  string
	SendToPhoneNumber string
}

type Context struct {
	MessageId string `json:"message_id,omitempty"`
}

// Base API Payload to send messages
type BaseMessagePayload struct {
	Context          *Context    `json:"context,omitempty"`
	To               string      `json:"to"`
	Type             MessageType `json:"type"`
	MessagingProduct string      `json:"messaging_product"`
	RecipientType    string      `json:"recipient_type"`
}

func NewBaseMessagePayload(to string, messageType MessageType) BaseMessagePayload {
	return BaseMessagePayload{
		To:               to,
		Type:             messageType,
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
	}
}

type BaseMessage interface {
	ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error)
}
