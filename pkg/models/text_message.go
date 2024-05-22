package models

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type TextMessage struct {
	Text         string
	AllowPreview bool
}

type TextMessageConfigs struct {
	Text         string `json:"text" validate:"required"`
	AllowPreview bool   `json:"allowPreview,omitempty"`
}

type TextMessageApiPayloadText struct {
	Body         string `json:"body" validate:"required"`
	AllowPreview bool   `json:"preview_url,omitempty"`
}

func (m *TextMessage) SetText(text string) {
	m.Text = text
}

func NewTextMessage(configs TextMessageConfigs) (*TextMessage, error) {
	err := utils.GetValidator().Struct(configs)
	if err != nil {
		return nil, fmt.Errorf("error validating text message config: %v", err)
	}
	return &TextMessage{
		Text:         configs.Text,
		AllowPreview: configs.AllowPreview,
	}, nil
}

// This function convert the TextMessage struct to WhatsApp API compatible JSON
func (m *TextMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := TextMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, "text"),
		Text: TextMessageApiPayloadText{
			Body:         m.Text,
			AllowPreview: m.AllowPreview,
		},
	}

	if configs.ReplyToMessageId != "" {
		jsonData.Context = &Context{
			MessageId: configs.ReplyToMessageId,
		}
	}

	jsonToReturn, err := json.Marshal(jsonData)

	if err != nil {
		return nil, fmt.Errorf("error marshalling json: %v", err)
	}

	return jsonToReturn, nil
}
