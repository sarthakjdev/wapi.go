package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal"
)

// ProductMessage represents a product message.
type ProductMessage struct {
}

// ProductMessageParams represents the parameters for creating a product message.
type ProductMessageParams struct {
}

// ProductMessageApiPayload represents the API payload for a product message.
type ProductMessageApiPayload struct {
	BaseMessagePayload
	Interactive ProductMessage `json:"interactive" validate:"required"`
}

// NewProductMessage creates a new product message with the given parameters.
func NewProductMessage(params ProductMessageParams) (*ProductMessage, error) {
	if err := internal.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}
	return &ProductMessage{}, nil
}

// ToJson converts the product message to JSON with the given configurations.
func (m *ProductMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := internal.GetValidator().Struct(configs); err != nil {
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
