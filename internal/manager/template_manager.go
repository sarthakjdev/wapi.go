package manager

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sarthakjdev/wapi.go/internal"
	"github.com/sarthakjdev/wapi.go/internal/request_client"
)

type MessageTemplateStatus string

const (
	MessageTemplateStatusApproved MessageTemplateStatus = "APPROVED"
	MessageTemplateStatusRejected MessageTemplateStatus = "REJECTED"
	MessageTemplateStatusPending  MessageTemplateStatus = "PENDING"
)

type MessageTemplateCategory string

const (
	MessageTemplateCategoryUtility        MessageTemplateCategory = "UTILITY"
	MessageTemplateCategoryMarketing      MessageTemplateCategory = "MARKETING"
	MessageTemplateCategoryAuthentication MessageTemplateCategory = "AUTHENTICATION"
)

type WhatsAppBusinessMessageTemplateNode struct {
	Id                         string                                                  `json:"id,omitempty"`
	Category                   MessageTemplateCategory                                 `json:"category,omitempty"`
	Components                 []WhatsAppBusinessHSMWhatsAppHSMComponent               `json:"components,omitempty"`
	CorrectCategory            string                                                  `json:"correct_category,omitempty"`
	CtaUrlLinkTrackingOptedOut bool                                                    `json:"cta_url_link_tracking_opted_out,omitempty"`
	Language                   string                                                  `json:"language,omitempty"`
	LibraryTemplateName        string                                                  `json:"library_template_name,omitempty"`
	MessageSendTtlSeconds      int                                                     `json:"message_send_ttl_seconds,omitempty"`
	Name                       string                                                  `json:"name,omitempty"`
	PreviousCategory           string                                                  `json:"previous_category,omitempty"`
	QualityScore               WhatsAppBusinessHSMWhatsAppBusinessHSMQualityScoreShape `json:"quality_score,omitempty"`
	RejectedReason             string                                                  `json:"rejected_reason,omitempty"`
	Status                     MessageTemplateStatus                                   `json:"status,omitempty"`
}

type TemplateManager struct {
	businessAccountId string
	apiAccessToken    string
	requester         *request_client.RequestClient
}

type TemplateManagerConfig struct {
	BusinessAccountId string
	ApiAccessToken    string
	Requester         *request_client.RequestClient
}

func NewTemplateManager(config *TemplateManagerConfig) *TemplateManager {
	return &TemplateManager{
		apiAccessToken:    config.ApiAccessToken,
		businessAccountId: config.BusinessAccountId,
		requester:         config.Requester,
	}
}

type WhatsAppBusinessTemplacesFetchResponseEdge struct {
	Data   []WhatsAppBusinessMessageTemplateNode      `json:"data,omitempty"`
	Paging internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
}

type WhatsappMessageTemplateEdge struct {
	Data    []WhatsAppBusinessMessageTemplateNode      `json:"data,omitempty"`
	Paging  internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
	Summary string                                     `json:"summary,omitempty"`
}

type WhatsAppBusinessHSMWhatsAppHSMComponentCard struct {
}

type WhatsAppBusinessHSMWhatsAppHSMComponentButton struct {
}

type WhatsAppBusinessHSMWhatsAppHSMComponentExample struct {
}

type WhatsAppBusinessHSMWhatsAppLimitedTimeOfferParameterShape struct {
}

type WhatsAppBusinessHSMWhatsAppHSMComponent struct {
	AddSecurityRecommendation bool                                                      `json:"add_security_recommendation,omitempty"`
	Buttons                   []WhatsAppBusinessHSMWhatsAppHSMComponentButton           `json:"buttons,omitempty"`
	Cards                     []WhatsAppBusinessHSMWhatsAppHSMComponentCard             `json:"cards,omitempty"`
	CodeExpirationMinutes     int                                                       `json:"code_expiration_minutes,omitempty"`
	Example                   WhatsAppBusinessHSMWhatsAppHSMComponentExample            `json:"example,omitempty"`
	Format                    string                                                    `json:"format,omitempty"`
	LimitedTimeOffer          WhatsAppBusinessHSMWhatsAppLimitedTimeOfferParameterShape `json:"limited_time_offer,omitempty"`
	Text                      string                                                    `json:"text,omitempty"`
	Type                      string                                                    `json:"type,omitempty"`
}

type WhatsAppBusinessHSMWhatsAppBusinessHSMQualityScoreShape struct {
	Date    int      `json:"date,omitempty"`
	Reasons []string `json:"reasons,omitempty"`
	Score   int      `json:"score,omitempty"`
}

func (manager *TemplateManager) FetchAll() (*WhatsAppBusinessTemplacesFetchResponseEdge, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{manager.businessAccountId, "/", "message_templates"}, ""), http.MethodGet)

	fields := []string{"id", "category", "components", "correct_category", "cta_url_link_tracking_opted_out", "language", "library_template_name", "message_send_ttl_seconds", "name", "previous_category", "quality_score", "rejected_reason", "status", "sub_category"}

	for _, field := range fields {
		apiRequest.AddField(request_client.BusinessApiRequestQueryParamField{
			Name:    field,
			Filters: map[string]string{},
		})
	}

	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}

	var response_to_return WhatsAppBusinessTemplacesFetchResponseEdge
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

func (manager *TemplateManager) Fetch(Id string) (*WhatsAppBusinessMessageTemplateNode, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{Id}, ""), http.MethodGet)
	fields := []string{"id", "category", "components", "correct_category", "cta_url_link_tracking_opted_out", "language", "library_template_name", "message_send_ttl_seconds", "name", "previous_category", "quality_score", "rejected_reason", "status", "sub_category"}
	for _, field := range fields {
		apiRequest.AddField(request_client.BusinessApiRequestQueryParamField{
			Name:    field,
			Filters: map[string]string{},
		})
	}
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}
	var response_to_return WhatsAppBusinessMessageTemplateNode
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

// this is button structure for the message template creation request
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

type MessageTemplateComponentType string

const (
	MessageTemplateComponentTypeGreeting         MessageTemplateComponentType = "GREETING"
	MessageTemplateComponentTypeHeader           MessageTemplateComponentType = "HEADER"
	MessageTemplateComponentTypeBody             MessageTemplateComponentType = "BODY"
	MessageTemplateComponentTypeFooter           MessageTemplateComponentType = "FOOTER"
	MessageTemplateComponentTypeButtons          MessageTemplateComponentType = "BUTTONS"
	MessageTemplateComponentTypeCarousel         MessageTemplateComponentType = "CAROUSEL"
	MessageTemplateComponentTypeLimitedTimeOffer MessageTemplateComponentType = "LIMITED_TIME_OFFER"
)

type MessageTemplateComponentFormat string

const (
	MessageTemplateComponentFormatText     MessageTemplateComponentFormat = "TEXT"
	MessageTemplateComponentFormatImage    MessageTemplateComponentFormat = "IMAGE"
	MessageTemplateComponentFormatDocument MessageTemplateComponentFormat = "DOCUMENT"
	MessageTemplateComponentFormatVideo    MessageTemplateComponentFormat = "VIDEO"
	MessageTemplateComponentFormatLocation MessageTemplateComponentFormat = "LOCATION"
)

// thus is component structure for the message template creation request
type WhatsappMessageTemplateComponentCreateOrUpdateRequestBody struct {
	Type    MessageTemplateComponentType                     `json:"type,omitempty"`
	Format  MessageTemplateComponentFormat                   `json:"format,omitempty"`
	Text    string                                           `json:"text,omitempty"`
	Buttons []WhatsappMessageTemplateButtonCreateRequestBody `json:"buttons,omitempty"`
}

func (component *WhatsappMessageTemplateComponentCreateOrUpdateRequestBody) AddButton(button WhatsappMessageTemplateButtonCreateRequestBody) {
	component.Buttons = append(component.Buttons, button)
}

// this is the request body for the message template creation request
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
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{manager.businessAccountId, "/", "message_templates"}, ""), http.MethodPost)
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	apiRequest.SetBody(string(jsonBody))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var response_to_return MessageTemplateCreationResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil

}

// this is the request body for the message template update request
type WhatsAppBusinessAccountMessageTemplateUpdateRequestBody struct {
	Components            []WhatsappMessageTemplateComponentCreateOrUpdateRequestBody `json:"components,omitempty"`
	Category              string                                                      `json:"category,omitempty"`
	MessageSendTtlSeconds int                                                         `json:"message_send_ttl_seconds,omitempty"`
}

func (manager *TemplateManager) Update(templateId string, updates WhatsAppBusinessAccountMessageTemplateUpdateRequestBody) (*MessageTemplateCreationResponse, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{templateId}, ""), http.MethodPost)
	jsonBody, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}
	apiRequest.SetBody(string(jsonBody))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var response_to_return MessageTemplateCreationResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil

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

func (tm *TemplateManager) FetchMessageTemplatePreviews() {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_template_previews/
}

type TemplateAnalyticsType struct {
}

type TemplatePerformanceAnalytics struct {
}

func (manager *TemplateManager) FetchAnalytics() (string, error) {
	// https://graph.facebook.com/LATEST-VERSION/WHATSAPP-BUSINESS-ACCOUNT-ID?fields=analytics&access_token=ACCESS-TOKEN

	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{}, "/"), http.MethodGet)

	response, err := apiRequest.Execute()

	return response, err

}

func (manager *TemplateManager) FetchPerformanceAnalytics(templateName, templateId string) (string, error) {
	// /v20.0/{whats-app-business-account-id}/template_performance_metrics
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/template_performance_metrics/
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{manager.businessAccountId, "template_performance_metrics"}, "/"), http.MethodGet)
	apiRequest.AddQueryParam("name", templateName)
	apiRequest.AddQueryParam("template_id", templateId)
	response, err := apiRequest.Execute()

	if err != nil {
		return "", err
	}

	return response, nil
}

func (manager *TemplateManager) MigrateFromOtherBusinessAccount(sourcePageNumber int, sourceWabaId int) (string, error) {
	// /{whats_app_business_account_id}/migrate_message_templates

	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{manager.businessAccountId, "migrate_message_templates"}, "/"), http.MethodGet)
	apiRequest.AddQueryParam("page_number", string(sourcePageNumber))
	apiRequest.AddQueryParam("source_waba_id", string(sourceWabaId))
	response, err := apiRequest.Execute()

	if err != nil {
		return "", err
	}

	return response, nil

	// Struct {
	// 	migrated_templates: List [
	// 	string
	// 	],
	// 	failed_templates: Map {
	// 	string: string
	// 	},
	// 	}

	// return type
}
