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
	AllowPreview bool   `json:"allowPreview"`
}

func (m *TextMessage) SetText(text string) {
	m.Text = text
}

func NewTextMessage(configs TextMessageConfigs) (*TextMessage, error) {
	err := utils.GetValidator().Struct(configs)
	fmt.Println("error validating text message config", err)
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
	// validate the configs
	err := utils.GetValidator().Struct(configs)
	if err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	data := make(map[string]interface{})
	if configs.ReplyToMessageId != "" {
		// add the context key to the json
		data["context"] = map[string]interface{}{
			"message_id": configs.ReplyToMessageId,
		}
	}

	data["type"] = "text"
	data["messaging_product"] = "whatsapp"
	data["recipient_type"] = "individual"
	data["text"] = map[string]interface{}{
		"body":        m.Text,
		"preview_url": m.AllowPreview,
	}
	data["to"] = configs.SendToPhoneNumber

	jsonToReturn, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("error marshalling json: %v", err)
	}

	return jsonToReturn, nil
}
