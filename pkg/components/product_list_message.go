package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal"
)

type Product struct {
	RetailerId string `json:"retailerId" validate:"required"`
}

func (p *Product) SetRetailerId(id string) {
	p.RetailerId = id
}

type ProductSection struct {
	Title    string    `json:"title" validate:"required"`
	Products []Product `json:"products" validate:"required"`
}

func (ps *ProductSection) SetTitle(title string) {
	ps.Title = title
}

func (ps *ProductSection) AddProduct(product Product) {
	ps.Products = append(ps.Products, product)
}

type ProductListMessageAction struct {
	Sections          []ProductSection `json:"sections" validate:"required"`
	CatalogId         string           `json:"catalogId" validate:"required"`
	ProductRetailerId string           `json:"productRetailerId" validate:"required"`
}

func (a *ProductListMessageAction) AddSection(section ProductSection) {
	a.Sections = append(a.Sections, section)
}

type ProductListMessageBody struct {
	Text string `json:"text" validate:"required"`
}

// ProductListMessage represents a product list message.
type ProductListMessage struct {
	Action ProductListMessageAction `json:"action" validate:"required"`
	Body   ProductListMessageBody   `json:"body" validate:"required"`
	Type   InteractiveMessageType   `json:"type" validate:"required"`
}

func (message *ProductListMessage) AddSection(section ProductSection) {
	message.Action.Sections = append(message.Action.Sections, section)
}

// ProductListMessageParams represents the parameters for creating a product list message.
type ProductListMessageParams struct {
	CatalogId         string `validate:"required"`
	ProductRetailerId string `validate:"required"`
	BodyText          string `validate:"required"`
	Sections          []ProductSection
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

	return &ProductListMessage{
		Type: InteractiveMessageTypeProductList,
		Body: ProductListMessageBody{
			Text: params.BodyText,
		},
		Action: ProductListMessageAction{
			CatalogId:         params.CatalogId,
			ProductRetailerId: params.ProductRetailerId,
			Sections:          params.Sections,
		},
	}, nil
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
