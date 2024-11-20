package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal"
)

type TemplateMessageComponentType string

const (
	TemplateMessageComponentTypeHeader TemplateMessageComponentType = "header"
	TemplateMessageComponentTypeBody   TemplateMessageComponentType = "body"
	TemplateMessageComponentTypeButton TemplateMessageComponentType = "button"
)

type TemplateMessageButtonComponentType string

const (
	TemplateMessageButtonComponentTypeQuickReply TemplateMessageButtonComponentType = "quick_reply"
	TemplateMessageButtonComponentTypeUrl        TemplateMessageButtonComponentType = "url"
	TemplateMessageButtonComponentTypeCatalog    TemplateMessageButtonComponentType = "catalog"
)

type TemplateMessageComponent interface {
	GetComponentType() string
}

type TemplateMessageComponentButtonType struct {
	Type       TemplateMessageComponentType       `json:"type" validate:"required"`       // Type of the template message component.
	SubType    TemplateMessageButtonComponentType `json:"sub_type" validate:"required"`   // Subtype of the template message component.
	Index      int                                `json:"index" validate:"required"`      // Index of the template message component.
	Parameters []TemplateMessageParameter         `json:"parameters" validate:"required"` // Parameters of the template message component.
}

func (t TemplateMessageComponentButtonType) GetComponentType() string {
	return string(t.Type)
}

type TemplateMessageComponentHeaderType struct {
	Type       TemplateMessageComponentType `json:"type" validate:"required"`       // Type of the template message component.
	Parameters []TemplateMessageParameter   `json:"parameters" validate:"required"` // Parameters of the template message component.
}

func (t TemplateMessageComponentHeaderType) GetComponentType() string {
	return string(t.Type)
}

type TemplateMessageComponentBodyType struct {
	Type       TemplateMessageComponentType `json:"type" validate:"required"`       // Type of the template message component.
	Parameters []TemplateMessageParameter   `json:"parameters" validate:"required"` // Parameters of the template message component.
}

func (t TemplateMessageComponentBodyType) GetComponentType() string {
	return string(t.Type)
}

type TemplateMessageParameterType string

const (
	TemplateMessageParameterTypeCurrency TemplateMessageParameterType = "currency"
	TemplateMessageParameterTypeDateTime TemplateMessageParameterType = "date_time"
	TemplateMessageParameterTypeDocument TemplateMessageParameterType = "document"
	TemplateMessageParameterTypeImage    TemplateMessageParameterType = "image"
	TemplateMessageParameterTypeText     TemplateMessageParameterType = "text"
	TemplateMessageParameterTypeVideo    TemplateMessageParameterType = "video"
)

type TemplateMessageParameterCurrency struct {
	FallbackValue string `json:"fallback_value" validate:"required"` // Fallback value of the currency parameter.
	Code          string `json:"code" validate:"required"`           // Code of the currency parameter.
	Amount1000    int    `json:"amount_1000" validate:"required"`    // Amount of the currency parameter.
}

type TemplateMessageParameterDateTime struct {
	FallbackValue string `json:"fallback_value" validate:"required"` // Fallback value of the date time parameter.
}

type TemplateMessageParameterDocument struct {
	Document interface{} `json:"document" validate:"required"` // Document of the document parameter.
}

type TemplateMessageParameterImage struct {
	Image interface{} `json:"image" validate:"required"` // Image of the image parameter.
}

type TemplateMessageParameterText struct {
	Text string `json:"text" validate:"required"` // Text of the text parameter.
}

type TemplateMessageParameterVideo struct {
	Video interface{} `json:"video" validate:"required"` // Video of the video parameter.
}

type TemplateMessageParameter interface {
	GetParameterType() string
}

type TemplateMessageBodyAndHeaderParameter struct {
	Type     TemplateMessageParameterType      `json:"type" validate:"required"` // Type of the template message parameter.
	Currency *TemplateMessageParameterCurrency `json:"currency,omitempty"`       // Currency of the template message parameter.
	DateTime *TemplateMessageParameterDateTime `json:"date_time,omitempty"`      // Date time of the template message parameter.
	Document *TemplateMessageParameterDocument `json:"document,omitempty"`       // Document of the template message parameter.
	Image    *TemplateMessageParameterImage    `json:"image,omitempty"`          // Image of the template message parameter.
	Text     *TemplateMessageParameterText     `json:"text,omitempty"`           // Text of the template message parameter.
	Video    *TemplateMessageParameterVideo    `json:"video,omitempty"`          // Video of the template message parameter.
}

func (t TemplateMessageBodyAndHeaderParameter) GetParameterType() string {
	return string(t.Type)
}

type TemplateMessageButtonParameterType string

const (
	TemplateMessageButtonParameterTypePayload TemplateMessageButtonParameterType = "payload"
	TemplateMessageButtonParameterTypeText    TemplateMessageButtonParameterType = "text"
)

type TemplateMessageButtonParameter struct {
	Type    TemplateMessageButtonParameterType `json:"type" validate:"required"` // Type of the template message parameter.
	Payload string                             `json:"payload,omitempty"`        // Payload of the template message parameter, required for type of button sub_type -> quick_reply
	Text    string                             `json:"text,omitempty"`           // Text of the template message parameter, required for type of button sub_type -> URL.
}

func (t TemplateMessageButtonParameter) GetParameterType() string {
	return string(t.Type)
}

type TemplateMessageLanguage struct {
	Code   string `json:"code" validate:"required"`
	Policy string `json:"policy" validate:"required"`
}

// TemplateMessage represents a template message.
type TemplateMessage struct {
	Name       string                     `json:"name" validate:"required"` // Name of the template message.
	Language   TemplateMessageLanguage    `json:"language" validate:"required"`
	Components []TemplateMessageComponent `json:"components" validate:"required"` // Components of the template message.
}

// TemplateMessageApiPayload represents the API payload for a template message.
type TemplateMessageApiPayload struct {
	BaseMessagePayload
	Template TemplateMessage `json:"template" validate:"required"`
}

// TemplateMessageConfigs represents the configurations for a template message.
type TemplateMessageConfigs struct {
	Name     string `json:"name" validate:"required"`     // Name of the template message.
	Language string `json:"language" validate:"required"` // Language of the template message.
}

// NewTemplateMessage creates a new instance of TemplateMessage.
func NewTemplateMessage(params *TemplateMessageConfigs) (*TemplateMessage, error) {
	return &TemplateMessage{
		Name: params.Name,
		Language: TemplateMessageLanguage{
			Code:   params.Language,
			Policy: "deterministic",
		},
	}, nil
}

// AddHeader adds a header to the template message.
// Template Messages in whatsapp can only have one header, so this function will override the last header added if any.
func (tm *TemplateMessage) AddHeader(params TemplateMessageComponentHeaderType) {
	var existingHeaderIndex int
	var found bool

	for i, component := range tm.Components {
		if TemplateMessageComponentType(component.GetComponentType()) == TemplateMessageComponentTypeHeader {
			existingHeaderIndex = i
			found = true
			break
		}
	}

	if found {
		// Override the existing header.
		tm.Components[existingHeaderIndex] = params
	} else {
		// Add the new header if no existing header is found.
		tm.Components = append(tm.Components, params)
	}
}

// AddBody adds a body to the template message.
// Template Messages in whatsapp can only have one body component, so this function will override the last body added if any.
func (tm *TemplateMessage) AddBody(params TemplateMessageComponentBodyType) {
	var existingBodyIndex int
	var found bool

	for i, component := range tm.Components {
		if TemplateMessageComponentType(component.GetComponentType()) == TemplateMessageComponentTypeBody {
			existingBodyIndex = i
			found = true
			break
		}
	}

	if found {
		// Override the existing body.
		tm.Components[existingBodyIndex] = params
	} else {
		// Add the new body if no existing body is found.
		tm.Components = append(tm.Components, params)
	}
}

func (tm *TemplateMessage) AddButton(params TemplateMessageComponentButtonType) error {
	// at max 10 buttons can be added to a template message
	numberOfButtons := 0
	for _, component := range tm.Components {
		if TemplateMessageComponentType(component.GetComponentType()) == TemplateMessageComponentTypeButton {
			numberOfButtons++
		}
	}

	if numberOfButtons >= 10 {
		return fmt.Errorf("maximum number of buttons reached")
	}
	tm.Components = append(tm.Components, params)
	return nil
}

// ToJson converts the template message to JSON.
func (m *TemplateMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := internal.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := TemplateMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeTemplate),
		Template:           *m,
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
