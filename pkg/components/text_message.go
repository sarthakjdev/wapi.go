package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type textMessage struct {
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

func (m *textMessage) SetText(text string) {
	m.Text = text
}

type TextMessageApiPayload struct {
	BaseMessagePayload `json:",inline"`
	Text               TextMessageApiPayloadText `json:"text" validate:"required"`
}

func NewTextMessage(configs TextMessageConfigs) (*textMessage, error) {
	err := utils.GetValidator().Struct(configs)
	if err != nil {
		return nil, fmt.Errorf("error validating text message config: %v", err)
	}
	return &textMessage{
		Text:         configs.Text,
		AllowPreview: configs.AllowPreview,
	}, nil
}

// This function convert the TextMessage struct to WhatsApp API compatible JSON
func (m *textMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := TextMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeText),
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
