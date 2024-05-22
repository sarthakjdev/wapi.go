package models

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type TextMessage struct {
	Text         string `json:"text" validate:"required"`
	AllowPreview bool   `json:"allowPreview"`
}

type TextMessageConfigs struct {
	Text         string
	AllowPreview bool
}

func (m *TextMessage) SetText(text string) {
	m.Text = text
}

func NewTextMessage(configs TextMessageConfigs) *TextMessage {
	err := utils.GetValidator().Struct(configs)
	fmt.Println("error validating text message config", err)
	if err != nil {
		return nil
	}
	return &TextMessage{
		Text:         configs.Text,
		AllowPreview: configs.AllowPreview,
	}
}

func (m *TextMessage) ToJson() ([]byte, error) {
	messageJson, err := json.Marshal(m)
	if err != nil {
		// emit a error event here
		return nil, err
	}
	fmt.Println("text message json is", string(messageJson))

	return messageJson, nil
}
