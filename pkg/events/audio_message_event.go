package events

import (
	"github.com/sarthakjdev/wapi.go/pkg/components"
)

type AudioMessageEvent struct {
	BaseMessageEvent
	Audio    components.AudioMessage
	MimeType string `json:"mimetype"`
	SHA256   string `json:"sha256"`
	MediaId  string `json:"media_id"`
}

func NewAudioMessageEvent(baseMessageEvent BaseMessageEvent, mediaId string, audio components.AudioMessage, mime_type, sha256 string) *AudioMessageEvent {
	return &AudioMessageEvent{
		BaseMessageEvent: baseMessageEvent,
		Audio:            audio,
		MimeType:         mime_type,
		SHA256:           sha256,
		MediaId:          mediaId,
	}
}
