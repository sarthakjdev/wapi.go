package components

// MessageType represents the type of message.
type MessageType string

// Constants for different message types.
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

// ApiCompatibleJsonConverterConfigs represents the configuration for converting to JSON.
type ApiCompatibleJsonConverterConfigs struct {
	ReplyToMessageId  string
	SendToPhoneNumber string
}

// Context represents the context of the message.
type Context struct {
	MessageId string `json:"message_id,omitempty"`
}

// BaseMessagePayload represents the base payload to send messages.
type BaseMessagePayload struct {
	Context          *Context    `json:"context,omitempty"`
	To               string      `json:"to"`
	Type             MessageType `json:"type"`
	MessagingProduct string      `json:"messaging_product"`
	RecipientType    string      `json:"recipient_type"`
}

// NewBaseMessagePayload creates a new instance of BaseMessagePayload.
func NewBaseMessagePayload(to string, messageType MessageType) BaseMessagePayload {
	return BaseMessagePayload{
		To:               to,
		Type:             messageType,
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
	}
}

// BaseMessage is an interface for sending messages.
type BaseMessage interface {
	ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error)
}
