package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type ProductMessage struct {
}

type ProductMessageParams struct {
}

type ProductMessageApiPayload struct {
	BaseMessagePayload
	Interactive ProductMessage `json:"interactive" validate:"required"`
}

func NewProductMessage(params ProductMessageParams) (*ProductMessage, error) {
	if err := utils.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}
	return &ProductMessage{}, nil
}

func (m *ProductMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := ProductMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeInteractive),
		Interactive:        *m,
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
