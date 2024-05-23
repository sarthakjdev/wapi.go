package components

type MessageType string

const (
	LocationMessageType    MessageType = "location"
	AudioMessageType       MessageType = "audio"
	VideoMessageType       MessageType = "video"
	DocumentMessageType    MessageType = "document"
	TextMessageType        MessageType = "text"
	ContactMessageType     MessageType = "contacts"
	InteractiveMessageType MessageType = "interactive"
	TemplateMessageType    MessageType = "template"
	ReactionMessageType    MessageType = "reaction"
	StickerMessageType     MessageType = "sticker"
	ImageMessageType       MessageType = "image"
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
