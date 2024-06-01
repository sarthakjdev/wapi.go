package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

// DocumentMessage represents a document message.
type DocumentMessage struct {
}

// DocumentMessageApiPayload represents the API payload for a document message.
type DocumentMessageApiPayload struct {
	BaseMessagePayload
	Document DocumentMessage `json:"document" validate:"required"`
}

// DocumentMessageConfigs represents the configurations for a document message.
type DocumentMessageConfigs struct {
}

// NewDocumentMessage creates a new DocumentMessage instance.
func NewDocumentMessage(params DocumentMessageConfigs) (*DocumentMessage, error) {
	if err := utils.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	return &DocumentMessage{}, nil
}

// ToJson converts the DocumentMessage instance to JSON.
func (dm *DocumentMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := DocumentMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeDocument),
		Document:           *dm,
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
