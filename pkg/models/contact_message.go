package models

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type ContactMessage struct {
}

type ContactMessageConfigs struct {
}

func NewContactMessage(configs ContactMessageConfigs) *ContactMessage {
	err := utils.GetValidator().Struct(configs)
	fmt.Println("error validating text message config", err)
	if err != nil {
		return nil
	}
	return &ContactMessage{}
}

func (m *ContactMessage) ToJson() ([]byte, error) {
	messageJson, err := json.Marshal(m)
	if err != nil {
		// emit a error event here
		return nil, err
	}
	fmt.Println("text message json is", string(messageJson))

	return messageJson, nil
}
