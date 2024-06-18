package manager

import "github.com/sarthakjdev/wapi.go/internal"

type MessageTemplate struct {
}

type MessageTemplateParam struct {
}

type TemplateManager struct{}

type TemplateManagerConfig struct{}

func NewTemplateManager(config *TemplateManagerConfig) *TemplateManager {
	return &TemplateManager{}
}

func (tm *TemplateManager) FetchAll() {
	// ! TODO: call this API endpoint
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/
}

type WhatsAppBusinessHSMWhatsAppHSMComponentCardGet struct {
}

type WhatsAppBusinessHSMWhatsAppHSMComponentButtonGet struct {
}

type WhatsAppBusinessHSMWhatsAppHSMComponentExample struct {
}

type WhatsAppBusinessHSMWhatsAppLimitedTimeOfferParameterShape struct {
}

type WhatsAppBusinessHSMWhatsAppHSMComponentGet struct {
	AddSecurityRecommendation bool                                                      `json:"add_security_recommendation,omitempty"`
	Buttons                   []WhatsAppBusinessHSMWhatsAppHSMComponentButtonGet        `json:"buttons,omitempty"`
	Cards                     []WhatsAppBusinessHSMWhatsAppHSMComponentCardGet          `json:"cards,omitempty"`
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

type WhatsappMessageTemplate struct {
	Id                         string                                                  `json:"id,omitempty"`
	Category                   string                                                  `json:"category,omitempty"`
	Components                 []WhatsAppBusinessHSMWhatsAppHSMComponentGet            `json:"components,omitempty"`
	CorrectCategory            string                                                  `json:"correct_category,omitempty"`
	CtaUrlLinkTrackingOptedOut bool                                                    `json:"cta_url_link_tracking_opted_out,omitempty"`
	Language                   string                                                  `json:"language,omitempty"`
	LibraryTemplateName        string                                                  `json:"library_template_name,omitempty"`
	MessageSendTtlSeconds      int                                                     `json:"message_send_ttl_seconds,omitempty"`
	Name                       string                                                  `json:"name,omitempty"`
	PreviousCategory           string                                                  `json:"previous_category,omitempty"`
	QualityScore               WhatsAppBusinessHSMWhatsAppBusinessHSMQualityScoreShape `json:"quality_score,omitempty"`
	RejectedReason             string                                                  `json:"rejected_reason,omitempty"`
	Status                     string                                                  `json:"status,omitempty"`
	SubCategory                string                                                  `json:"sub_category,omitempty"`
}

type WhatsappMessageTemplateEdge struct {
	Data    []WhatsappMessageTemplate                  `json:"data,omitempty"`
	Paging  internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
	Summary string                                     `json:"summary,omitempty"`
}

func (tm *TemplateManager) Fetch(Id string) {
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-hsm/
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

// thus is component structure for the message template creation request
type WhatsappMessageTemplateComponentCreateRequestBody struct {

	// enum {GREETING, HEADER, BODY, FOOTER, BUTTONS, CAROUSEL, LIMITED_TIME_OFFER}
	Type string `json:"type,omitempty"`

	// enum {TEXT, IMAGE, DOCUMENT, VIDEO, LOCATION}
	Format  string                                           `json:"format,omitempty"`
	Text    string                                           `json:"text,omitempty"`
	Buttons []WhatsappMessageTemplateButtonCreateRequestBody `json:"buttons,omitempty"`
}

// this is the request body for the message template creation request
type WhatsappMessageTemplateCreateRequestBody struct {
	AllowCategoryChange bool `json:"allow_category_change,omitempty" `

	// enum {UTILITY, MARKETING, AUTHENTICATION}
	Category string `json:"category,omitempty" validate:"required"`

	Components                  []WhatsappMessageTemplateComponentCreateRequestBody `json:"components" validate:"required"`
	Name                        string                                              `json:"name,omitempty" validate:"required"`
	Language                    string                                              `json:"language" validate:"required"`
	LibraryTemplateName         string                                              `json:"library_template_name,omitempty"`
	LibraryTemplateButtonInputs []WhatsappMessageTemplateButtonCreateRequestBody    `json:"library_template_button_inputs,omitempty"`
}

func (tm *TemplateManager) Create() {

	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/message_templates/#:~:text=You%20can%20make%20a%20POST%20request%20to%20message_templates%20edge%20from%20the%20following%20paths%3A
}

// this is the request body for the message template update request
type WhatsAppBusinessAccountMessageTemplateUpdateRequestBody struct {
	Components            []WhatsappMessageTemplateComponentCreateRequestBody `json:"components,omitempty"`
	Category              string                                              `json:"category,omitempty"`
	MessageSendTtlSeconds int                                                 `json:"message_send_ttl_seconds,omitempty"`
}

func (tm *TemplateManager) Update() {
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

func (bm *TemplateManager) FetchAnalytics() {
	// https://graph.facebook.com/LATEST-VERSION/WHATSAPP-BUSINESS-ACCOUNT-ID?fields=analytics&access_token=ACCESS-TOKEN

}
