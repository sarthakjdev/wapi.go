package models

type MessageType string

const (
	MessageTypeText        MessageType = "text"
	MessageTypeAudio       MessageType = "audio"
	MessageTypeDocument    MessageType = "document"
	MessageTypeContact     MessageType = "contact"
	MessageTypeLocation    MessageType = "location"
	MessageTypeImage       MessageType = "image"
	MessageTypeVideo       MessageType = "video"
	MessageTypeInteractive MessageType = "interactive"
	MessageTypeTemplate    MessageType = "template"
)

type BaseMessage interface {
	ToJson() ([]byte, error)
}
