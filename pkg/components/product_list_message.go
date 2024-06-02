package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal"
)

// ProductListMessage represents a product list message.
type ProductListMessage struct {
}

// ProductListMessageParams represents the parameters for creating a product list message.
type ProductListMessageParams struct {
}

// ProductListMessageApiPayload represents the API payload for a product list message.
type ProductListMessageApiPayload struct {
	BaseMessagePayload
	Interactive ProductListMessage `json:"interactive" validate:"required"`
}

// NewProductListMessage creates a new product list message.
func NewProductListMessage(params ProductListMessageParams) (*ProductListMessage, error) {
	if err := internal.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	return &ProductListMessage{}, nil
}

// ToJson converts the product list message to JSON.
func (m *ProductListMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := internal.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := ProductListMessageApiPayload{
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
