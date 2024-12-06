package components

import (
	"encoding/json"
	"fmt"

	"github.com/wapikit/wapi.go/internal"
)

type ProductMessageBody struct {
	Text string `json:"text" validate:"required"`
}

type productMessageAction struct {
	CatalogId         string `json:"catalog_id" validate:"required"`
	ProductRetailerId string `json:"product_retailer_id" validate:"required"`
}

// ProductMessage represents a product message.
type ProductMessage struct {
	Type   InteractiveMessageType `json:"type" validate:"required"`
	Body   ProductMessageBody     `json:"body" validate:"required"`
	Action productMessageAction   `json:"action" validate:"required"`
}

// ProductMessageParams represents the parameters for creating a product message.
type ProductMessageParams struct {
	CatalogId         string `validate:"required"`
	ProductRetailerId string `validate:"required"`
	BodyText          string `validate:"required"`
}

// ProductMessageApiPayload represents the API payload for a product message.
type ProductMessageApiPayload struct {
	BaseMessagePayload
	Interactive ProductMessage `json:"interactive" validate:"required"`
}

func (m *ProductMessage) SetHeader() {
}

func (m *ProductMessage) SetBodyText(text string) {
	m.Body.Text = text
}

func (m *ProductMessage) SetCatalogId(id string) {
	m.Action.CatalogId = id
}

func (m *ProductMessage) SetProductRetailerId(id string) {
	m.Action.ProductRetailerId = id
}

// NewProductMessage creates a new product message with the given parameters.
func NewProductMessage(params ProductMessageParams) (*ProductMessage, error) {
	if err := internal.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}
	return &ProductMessage{
		Type: InteractiveMessageTypeProduct,
		Body: ProductMessageBody{
			Text: params.BodyText,
		},
		Action: productMessageAction{
			CatalogId:         params.CatalogId,
			ProductRetailerId: params.ProductRetailerId,
		},
	}, nil
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
