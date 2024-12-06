package manager

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/wapikit/wapi.go/internal"
	"github.com/wapikit/wapi.go/internal/request_client"
)

type MessageTemplateType string

const (
	MessageTemplateTypeHeader MessageTemplateType = "HEADER"
	MessageTemplateTypeBody   MessageTemplateType = "BODY"
	MessageTemplateTypeButton MessageTemplateType = "BUTTONS"
	MessageTemplateTypeFooter MessageTemplateType = "FOOTER"
)

// MessageTemplateStatus represents the status of a WhatsApp Business message template.
type MessageTemplateStatus string

// Constants representing different message template statuses.
const (
	MessageTemplateStatusApproved MessageTemplateStatus = "APPROVED"
	MessageTemplateStatusRejected MessageTemplateStatus = "REJECTED"
	MessageTemplateStatusPending  MessageTemplateStatus = "PENDING"
)

// MessageTemplateCategory represents the category of a WhatsApp Business message template.
type MessageTemplateCategory string

// Constants representing different message template categories.
const (
	MessageTemplateCategoryUtility        MessageTemplateCategory = "UTILITY"
	MessageTemplateCategoryMarketing      MessageTemplateCategory = "MARKETING"
	MessageTemplateCategoryAuthentication MessageTemplateCategory = "AUTHENTICATION"
)

// WhatsAppBusinessMessageTemplateNode represents a WhatsApp Business message template.
type WhatsAppBusinessMessageTemplateNode struct {
	Id                         string                                    `json:"id,omitempty"`
	Category                   MessageTemplateCategory                   `json:"category,omitempty"`
	Components                 []WhatsAppBusinessHSMWhatsAppHSMComponent `json:"components,omitempty"`
	CorrectCategory            string                                    `json:"correct_category,omitempty"`
	CtaUrlLinkTrackingOptedOut bool                                      `json:"cta_url_link_tracking_opted_out,omitempty"`
	Language                   string                                    `json:"language,omitempty"`
	LibraryTemplateName        string                                    `json:"library_template_name,omitempty"`
	MessageSendTtlSeconds      int                                       `json:"message_send_ttl_seconds,omitempty"`
	Name                       string                                    `json:"name,omitempty"`
	PreviousCategory           string                                    `json:"previous_category,omitempty"`
	QualityScore               TemplateMessageQualityScore               `json:"quality_score,omitempty"`
	RejectedReason             string                                    `json:"rejected_reason,omitempty"`
	Status                     MessageTemplateStatus                     `json:"status,omitempty"`
}

// TemplateManager is responsible for managing WhatsApp Business message templates.
type TemplateManager struct {
	businessAccountId string
	apiAccessToken    string
	requester         *request_client.RequestClient
}

// TemplateManagerConfig represents the configuration for creating a new TemplateManager.
type TemplateManagerConfig struct {
	BusinessAccountId string
	ApiAccessToken    string
	Requester         *request_client.RequestClient
}

// NewTemplateManager creates a new TemplateManager with the given configuration.
func NewTemplateManager(config *TemplateManagerConfig) *TemplateManager {
	return &TemplateManager{
		apiAccessToken:    config.ApiAccessToken,
		businessAccountId: config.BusinessAccountId,
		requester:         config.Requester,
	}
}

// WhatsAppBusinessTemplatesFetchResponseEdge represents the response structure for fetching WhatsApp Business message templates.
type WhatsAppBusinessTemplatesFetchResponseEdge struct {
	Data   []WhatsAppBusinessMessageTemplateNode      `json:"data,omitempty"`
	Paging internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
}

// TemplateMessageComponentCard represents a card component in a WhatsApp Business message template.
type TemplateMessageComponentCard struct {
}

type TemplateMessageButtonType string

const (
	TemplateMessageButtonTypeQuickReply  TemplateMessageButtonType = "QUICK_REPLY"
	TemplateMessageButtonTypeUrl         TemplateMessageButtonType = "URL"
	TemplateMessageButtonTypePhoneNumber TemplateMessageButtonType = "PHONE_NUMBER"
	TemplateMessageButtonTypeCopyCode    TemplateMessageButtonType = "COPY_CODE"
)

// TemplateMessageComponentButton represents a button component in a WhatsApp Business message template.
type TemplateMessageComponentButton struct {
	Type        TemplateMessageButtonType `json:"type,omitempty"`
	Text        string                    `json:"text,omitempty"`
	PhoneNumber string                    `json:"phone_number,omitempty"` // required when Type = PHONE_NUMBER
	Example     string                    `json:"example,omitempty"`      // required when Type = URL and button has a variable
	Url         string                    `json:"url,omitempty"`          // required when Type = URL
}

// TemplateMessageComponentExample represents an example component in a WhatsApp Business message template.
type TemplateMessageComponentExample struct {
	HeaderHandle []string   `json:"header_handle,omitempty"`
	HeaderText   []string   `json:"header_text,omitempty"`
	BodyText     [][]string `json:"body_text,omitempty"`
}

type TemplateMessageComponentButtonExample []string

// TemplateMessageLimitedTimeOfferParameter represents a limited time offer parameter in a WhatsApp Business message template.
type TemplateMessageLimitedTimeOfferParameter struct {
}

// WhatsAppBusinessHSMWhatsAppHSMComponent represents a component in a WhatsApp Business message template.
type WhatsAppBusinessHSMWhatsAppHSMComponent struct {
	AddSecurityRecommendation bool                                     `json:"add_security_recommendation,omitempty"`
	Buttons                   []TemplateMessageComponentButton         `json:"buttons,omitempty"`
	Cards                     []TemplateMessageComponentCard           `json:"cards,omitempty"`
	CodeExpirationMinutes     int                                      `json:"code_expiration_minutes,omitempty"`
	Example                   TemplateMessageComponentExample          `json:"example,omitempty"`
	Format                    MessageTemplateComponentFormat           `json:"format,omitempty"`
	LimitedTimeOffer          TemplateMessageLimitedTimeOfferParameter `json:"limited_time_offer,omitempty"`
	Text                      string                                   `json:"text,omitempty"`
	Type                      MessageTemplateStatus                    `json:"type,omitempty"`
}

// TemplateMessageQualityScore represents the quality score of a WhatsApp Business message template.
type TemplateMessageQualityScore struct {
	Date    int      `json:"date,omitempty"`
	Reasons []string `json:"reasons,omitempty"`
	Score   int      `json:"score,omitempty"`
}

// FetchAll fetches all WhatsApp Business message templates.
func (manager *TemplateManager) FetchAll() (*WhatsAppBusinessTemplatesFetchResponseEdge, error) {
	apiRequest := manager.requester.NewApiRequest(strings.Join([]string{manager.businessAccountId, "/", "message_templates"}, ""), http.MethodGet)

	fields := []string{"id", "category", "components", "correct_category", "cta_url_link_tracking_opted_out", "language", "library_template_name", "message_send_ttl_seconds", "name", "previous_category", "quality_score", "rejected_reason", "status", "sub_category"}

	for _, field := range fields {
		apiRequest.AddField(request_client.ApiRequestQueryParamField{
			Name:    field,
			Filters: map[string]string{},
		})
	}

	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}

	var responseToReturn WhatsAppBusinessTemplatesFetchResponseEdge
	json.Unmarshal([]byte(response), &responseToReturn)
	return &responseToReturn, nil
}

// Fetch fetches a single WhatsApp Business message template by its ID.
func (manager *TemplateManager) Fetch(Id string) (*WhatsAppBusinessMessageTemplateNode, error) {
	apiRequest := manager.requester.NewApiRequest(strings.Join([]string{Id}, ""), http.MethodGet)
	fields := []string{"id", "category", "components", "correct_category", "cta_url_link_tracking_opted_out", "language", "library_template_name", "message_send_ttl_seconds", "name", "previous_category", "quality_score", "rejected_reason", "status", "sub_category"}
	for _, field := range fields {
		apiRequest.AddField(request_client.ApiRequestQueryParamField{
			Name:    field,
			Filters: map[string]string{},
		})
	}
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}
	var responseToReturn WhatsAppBusinessMessageTemplateNode
	json.Unmarshal([]byte(response), &responseToReturn)
	return &responseToReturn, nil
}

// WhatsappMessageTemplateButtonCreateRequestBody represents the request body for creating a button in a message template.
type WhatsappMessageTemplateButtonCreateRequestBody struct {
	// enum {QUICK_REPLY, URL, PHONE_NUMBER, OTP, MPM, CATALOG, FLOW, VOICE_CALL}
	Type                 string `json:"type,omitempty"`
	Text                 string `json:"text,omitempty"`
	Url                  string `json:"url,omitempty"`
	PhoneNumber          string `json:"phone_number,omitempty"`
	Example              string `json:"example,omitempty"`
	FlowId               string `json:"flow_id,omitempty"`
	ZeroTapTermsAccepted bool   `json:"zero_tap_terms_accepted,omitempty"`
}

// MessageTemplateComponentType represents the type of a message template component.
type MessageTemplateComponentType string

// Constants representing different message template component types.
const (
	MessageTemplateComponentTypeGreeting         MessageTemplateComponentType = "GREETING"
	MessageTemplateComponentTypeHeader           MessageTemplateComponentType = "HEADER"
	MessageTemplateComponentTypeBody             MessageTemplateComponentType = "BODY"
	MessageTemplateComponentTypeFooter           MessageTemplateComponentType = "FOOTER"
	MessageTemplateComponentTypeButtons          MessageTemplateComponentType = "BUTTONS"
	MessageTemplateComponentTypeCarousel         MessageTemplateComponentType = "CAROUSEL"
	MessageTemplateComponentTypeLimitedTimeOffer MessageTemplateComponentType = "LIMITED_TIME_OFFER"
)

// MessageTemplateComponentFormat represents the format of a message template component.
type MessageTemplateComponentFormat string

// Constants representing different message template component formats.
const (
	MessageTemplateComponentFormatText     MessageTemplateComponentFormat = "TEXT"
	MessageTemplateComponentFormatImage    MessageTemplateComponentFormat = "IMAGE"
	MessageTemplateComponentFormatDocument MessageTemplateComponentFormat = "DOCUMENT"
	MessageTemplateComponentFormatVideo    MessageTemplateComponentFormat = "VIDEO"
	MessageTemplateComponentFormatLocation MessageTemplateComponentFormat = "LOCATION"
)

// WhatsappMessageTemplateComponentCreateOrUpdateRequestBody represents the request body for creating or updating a component in a message template.
type WhatsappMessageTemplateComponentCreateOrUpdateRequestBody struct {
	Type    MessageTemplateComponentType                     `json:"type,omitempty"`
	Format  MessageTemplateComponentFormat                   `json:"format,omitempty"`
	Text    string                                           `json:"text,omitempty"`
	Buttons []WhatsappMessageTemplateButtonCreateRequestBody `json:"buttons,omitempty"`
}

// AddButton adds a button to the component.
func (component *WhatsappMessageTemplateComponentCreateOrUpdateRequestBody) AddButton(button WhatsappMessageTemplateButtonCreateRequestBody) {
	component.Buttons = append(component.Buttons, button)
}

// WhatsappMessageTemplateCreateRequestBody represents the request body for creating a message template.
type WhatsappMessageTemplateCreateRequestBody struct {
	AllowCategoryChange bool `json:"allow_category_change,omitempty" `

	// enum {UTILITY, MARKETING, AUTHENTICATION}
	Category string `json:"category,omitempty" validate:"required"`

	Components                  []WhatsappMessageTemplateComponentCreateOrUpdateRequestBody `json:"components" validate:"required"`
	Name                        string                                                      `json:"name,omitempty" validate:"required"`
	Language                    string                                                      `json:"language" validate:"required"`
	LibraryTemplateName         string                                                      `json:"library_template_name,omitempty"`
	LibraryTemplateButtonInputs []WhatsappMessageTemplateButtonCreateRequestBody            `json:"library_template_button_inputs,omitempty"`
}

func (body *WhatsappMessageTemplateCreateRequestBody) AddComponent(component WhatsappMessageTemplateComponentCreateOrUpdateRequestBody) {
	body.Components = append(body.Components, component)
}

type MessageTemplateCreationResponse struct {
	Id       string                  `json:"id,omitempty"`
	Status   MessageTemplateStatus   `json:"status,omitempty"`
	Category MessageTemplateCategory `json:"category,omitempty"`
}

func (manager *TemplateManager) Create(body WhatsappMessageTemplateCreateRequestBody) (*MessageTemplateCreationResponse, error) {
	apiRequest := manager.requester.NewApiRequest(strings.Join([]string{manager.businessAccountId, "/", "message_templates"}, ""), http.MethodPost)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	apiRequest.SetBody(string(jsonBody))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var responseToReturn MessageTemplateCreationResponse
	json.Unmarshal([]byte(response), &responseToReturn)
	return &responseToReturn, nil

}

// this is the request body for the message template update request
type WhatsAppBusinessAccountMessageTemplateUpdateRequestBody struct {
	Components            []WhatsappMessageTemplateComponentCreateOrUpdateRequestBody `json:"components,omitempty"`
	Category              string                                                      `json:"category,omitempty"`
	MessageSendTtlSeconds int                                                         `json:"message_send_ttl_seconds,omitempty"`
}

func (manager *TemplateManager) Update(templateId string, updates WhatsAppBusinessAccountMessageTemplateUpdateRequestBody) (*MessageTemplateCreationResponse, error) {
	apiRequest := manager.requester.NewApiRequest(strings.Join([]string{templateId}, ""), http.MethodPost)
	jsonBody, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}
	apiRequest.SetBody(string(jsonBody))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var responseToReturn MessageTemplateCreationResponse
	json.Unmarshal([]byte(response), &responseToReturn)
	return &responseToReturn, nil

	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-hsm/#:~:text=2.0%20Access%20Token-,Updating,-You%20can%20update
}

type WhatsAppBusinessAccountMessageTemplateDeleteRequestBody struct {
	HsmId string `json:"hsm_id,omitempty"`
	Name  string `json:"name,omitempty"`
}

func (tm *TemplateManager) Delete(id string) {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/#:~:text=on%20this%20endpoint.-,Deleting,-You%20can%20dissociate
}

type WhatsAppBusinessAccountMessageTemplatePreviewButton struct {
	AutoFillText string `json:"auto_fill_text,omitempty"`
	Text         string `json:"text,omitempty"`
}

type TemplateMessagePreviewNode struct {
	Body     string                                                `json:"body,omitempty"`
	Buttons  []WhatsAppBusinessAccountMessageTemplatePreviewButton `json:"buttons,omitempty"`
	Footer   string                                                `json:"footer,omitempty"`
	Header   string                                                `json:"header,omitempty"`
	Language string                                                `json:"language,omitempty"`
}

type TemplateMessagePreviewEdge struct {
	Data   []TemplateMessagePreviewNode               `json:"data,omitempty"`
	Paging internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
}

type TemplateMigrationResponse struct {
	MigratedTemplates []string          `json:"migrated_templates,omitempty"`
	FailedTemplates   map[string]string `json:"failed_templates,omitempty"`
}

func (manager *TemplateManager) MigrateFromOtherBusinessAccount(sourcePageNumber int, sourceWabaId int) (*TemplateMigrationResponse, error) {
	apiRequest := manager.requester.NewApiRequest(strings.Join([]string{manager.businessAccountId, "migrate_message_templates"}, "/"), http.MethodGet)
	apiRequest.AddQueryParam("page_number", string(sourcePageNumber))
	apiRequest.AddQueryParam("source_waba_id", string(sourceWabaId))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}
	var responseToReturn TemplateMigrationResponse
	json.Unmarshal([]byte(response), &responseToReturn)
	return &responseToReturn, nil
}
