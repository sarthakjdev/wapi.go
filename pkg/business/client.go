package business

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sarthakjdev/wapi.go/internal"
	"github.com/sarthakjdev/wapi.go/internal/manager"
	"github.com/sarthakjdev/wapi.go/internal/request_client"
)

type BusinessClient struct {
	BusinessAccountId string `json:"businessAccountId" validate:"required"`
	AccessToken       string `json:"accessToken" validate:"required"`
	PhoneNumber       *manager.PhoneNumberManager
	Template          *manager.TemplateManager
	requester         *request_client.RequestClient
}

type BusinessClientConfig struct {
	BusinessAccountId string `json:"businessAccountId" validate:"required"`
	AccessToken       string `json:"accessToken" validate:"required"`
	Requester         *request_client.RequestClient
}

func NewBusinessClient(config *BusinessClientConfig) *BusinessClient {
	return &BusinessClient{
		BusinessAccountId: config.BusinessAccountId,
		AccessToken:       config.AccessToken,
		PhoneNumber:       manager.NewPhoneNumberManager(&manager.PhoneNumberManagerConfig{}),
		Template:          manager.NewTemplateManager(&manager.TemplateManagerConfig{}),
		requester:         config.Requester,
	}
}

func (client *BusinessClient) newBusinessApiRequest(path, method string) *WhatsappBusinessManagementApiRequest {
	return &WhatsappBusinessManagementApiRequest{
		Path:        path,
		Fields:      []BusinessApiRequestQueryParamField{},
		Requester:   client.requester,
		Method:      method,
		QueryParams: map[string]string{},
	}
}

type BusinessApiRequestQueryParamField struct {
	Name    string
	Filters map[string]string
}

func (field *BusinessApiRequestQueryParamField) AddFilter(key, value string) {
	field.Filters[key] = value
}

type WhatsappBusinessManagementApiRequest struct {
	Path        string
	Method      string
	Body        string
	Fields      []BusinessApiRequestQueryParamField
	QueryParams map[string]string
	Requester   *request_client.RequestClient
}

func (request *WhatsappBusinessManagementApiRequest) AddField(field BusinessApiRequestQueryParamField) *BusinessApiRequestQueryParamField {
	// * NOTE:  when we say we need to add a field to the request, it means we need to add a query param to the request
	// * note that if there need to be multiple fields in a single request then the list of fields should be command separated
	// * for example: fields=field1,field2,field3
	// * also note that if there filters in a field then they should be called like a function in the param string, for ex: fields=field1.filter1(value1).filter2(value2),field2.filter1(value1)
	request.Fields = append(request.Fields, field)

	return &field
}

func (request *WhatsappBusinessManagementApiRequest) AddQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *WhatsappBusinessManagementApiRequest) SetMethod(method string) {
	request.Method = method
}

func (request *WhatsappBusinessManagementApiRequest) SetBody(body string) {
	request.Body = body
}

func (businessRequest *WhatsappBusinessManagementApiRequest) Execute() (string, error) {
	// check if there are any fields in the request
	var queryParam = map[string]string{}
	if len(businessRequest.Fields) > 0 {
		fieldsString := ""
		for _, field := range businessRequest.Fields {
			newFieldString := ""
			if fieldsString != "" {
				newFieldString = ","
			}
			filterString := ""
			for key, value := range field.Filters {
				filterString += strings.Join([]string{".", key, "(", value, ")"}, "")
			}
			newFieldString = strings.Join([]string{field.Name, filterString}, "")
			fieldsString += newFieldString
		}

		queryParam["fields"] = fieldsString
	}

	if len(businessRequest.QueryParams) > 0 {
		for key, value := range businessRequest.QueryParams {
			queryParam[key] = value
		}
	}

	fmt.Println("Query params are", queryParam)

	response, err := businessRequest.Requester.Request(request_client.RequestCloudApiParams{
		Path:       businessRequest.Path,
		Body:       businessRequest.Body,
		Method:     businessRequest.Method,
		QueryParam: queryParam,
	})

	if err != nil {
		fmt.Println("Error while executing business api request", err)
		return "", nil
	}

	fmt.Println("Response from business api is", response)
	return response, err
}

func (bc *BusinessClient) GetBusinessId() string {
	return bc.BusinessAccountId
}

func (bc *BusinessClient) SetBusinessId(id string) {
	bc.BusinessAccountId = id
}

type WhatsAppBusinessAccountRequestField string

const ()

type WhatsappBusinessAccount struct {
	BusinessVerificationStatus string `json:"business_verification_status,omitempty"`
	Country                    string `json:"country,omitempty"`
	Currency                   string `json:"currency,omitempty"`
	IsTemplateAnalyticsEnabled string `json:"is_enabled_for_insights,omitempty"`
	MessageTemplateNamespace   string `json:"message_template_namespace,omitempty"`
	Name                       string `json:"name,omitempty"`
	OwnershipType              string `json:"ownership_type,omitempty"`
	PrimaryFundingId           string `json:"primary_funding_id,omitempty"`
	PurchaseOrderNumber        string `json:"purchase_order_number,omitempty"`
	TimezoneId                 string `json:"timezone_id,omitempty"`
}

type MessageTemplatePreview struct {
}

type ProductCatalog struct {
}

type FetchBusinessAccountResponse struct {
	Id                       string `json:"id" validate:"required"`
	Name                     string `json:"name" validate:"required"`
	TimezoneId               string `json:"timezone_id" validate:"required"`
	MessageTemplateNamespace string `json:"message_template_namespace" validate:"required"`
}

func (client *BusinessClient) Fetch() FetchBusinessAccountResponse {
	apiRequest := client.newBusinessApiRequest(client.BusinessAccountId, http.MethodGet)
	response, err := apiRequest.Execute()
	if err != nil {
		// return wapi.go custom error here
		fmt.Println("Error while fetching business account", err)
	}
	var response_to_return FetchBusinessAccountResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return response_to_return
}

type AnalyticsRequestGranularityType string

const (
	AnalyticsRequestGranularityTypeHalfHour AnalyticsRequestGranularityType = "HALF_HOUR"
	AnalyticsRequestGranularityTypeDay      AnalyticsRequestGranularityType = "DAY"
	AnalyticsRequestGranularityTypeMonth    AnalyticsRequestGranularityType = "MONTH"
)

type WhatsAppBusinessAccountAnalyticsProductType int

const (
	WhatsAppBusinessAccountAnalyticsProductTypeNotificationMessages    WhatsAppBusinessAccountAnalyticsProductType = 0
	WhatsAppBusinessAccountAnalyticsProductTypeCustomerSupportMessages WhatsAppBusinessAccountAnalyticsProductType = 2
)

type WhatsappBusinessAccountAnalyticsOptions struct {
	Start        time.Time                       `json:"start" validate:"required"`
	End          time.Time                       `json:"end" validate:"required"`
	Granularity  AnalyticsRequestGranularityType `json:"granularity" validate:"required"`
	PhoneNumbers []string                        `json:"phone_numbers,omitempty"`

	// * NOT SUPPORTED AS OF NOW
	// ProductTypes []WhatsAppBusinessAccountAnalyticsProductType `json:"product_types,omitempty"`
	CountryCodes []string `json:"country_codes,omitempty"`
}

type AnalyticsDataPoint struct {
	Start     int `json:"start,omitempty"`
	End       int `json:"end,omitempty"`
	Sent      int `json:"sent,omitempty"`
	Delivered int `json:"delivered,omitempty"`
}

type WhatsappBusinessAccountAnalyticsResponse struct {
	PhoneNumbers []string             `json:"phone_numbers,omitempty"`
	Granularity  string               `json:"granularity,omitempty"`
	DataPoints   []AnalyticsDataPoint `json:"data_points,omitempty"`
}

func (client *BusinessClient) FetchAnalytics(options WhatsappBusinessAccountAnalyticsOptions) {
	apiRequest := client.newBusinessApiRequest(client.BusinessAccountId, http.MethodGet)
	analyticsField := apiRequest.AddField(BusinessApiRequestQueryParamField{
		Name:    "analytics",
		Filters: map[string]string{},
	})
	analyticsField.AddFilter("start", fmt.Sprint(options.Start.Unix()))
	analyticsField.AddFilter("end", fmt.Sprint(options.End.Unix()))
	analyticsField.AddFilter("granularity", string(options.Granularity))

	if len(options.PhoneNumbers) > 0 {
		// get specific phone numbers
		analyticsField.AddFilter("phone_numbers", strings.Join(options.PhoneNumbers, ","))
	} else {
		// get all phone numbers
		analyticsField.AddFilter("phone_numbers", "[]")
	}

	if len(options.CountryCodes) > 0 {
		analyticsField.AddFilter("country_codes", strings.Join(options.CountryCodes, ","))
	} else {
		// get all country codes
		analyticsField.AddFilter("country_codes", "[]")
	}
	response, err := apiRequest.Execute()
	if err != nil {
		// return wapi.go custom error here
		fmt.Println("Error while fetching business account", err)
	}
	var response_to_return WhatsappBusinessAccountAnalyticsResponse
	json.Unmarshal([]byte(response), &response_to_return)
	fmt.Println("Response to return is", response_to_return)
}

type ConversationCategoryType string

const (
	ConversationCategoryTypeAuthentication ConversationCategoryType = "AUTHENTICATION"
	ConversationCategoryTypeMarketing      ConversationCategoryType = "MARKETING"
	ConversationCategoryTypeService        ConversationCategoryType = "SERVICE"
	ConversationCategoryTypeUtility        ConversationCategoryType = "UTILITY"
)

type ConversationType string

const (
	ConversationTypeFreeEntry ConversationType = "FREE_ENTRY"
	ConversationTypeFreeTier  ConversationType = "FREE_TIER"
	ConversationTypeRegular   ConversationType = "REGULAR"
)

type ConversationDirection string

const (
	ConversationDirectionBusinessInitiated ConversationDirection = "BUSINESS_INITIATED"
	ConversationDirectionUserInitiated     ConversationDirection = "USER_INITIATED"
)

type ConversationDimensionType string

const (
	ConversationDimensionTypeConversationCategory  ConversationDimensionType = "CONVERSATION_CATEGORY"
	ConversationDimensionTypeConversationDirection ConversationDimensionType = "CONVERSATION_DIRECTION"
	ConversationDimensionTypeConversationType      ConversationDimensionType = "CONVERSATION_TYPE"
	ConversationDimensionTypeCountry               ConversationDimensionType = "COUNTRY"
	ConversationDimensionTypePhone                 ConversationDimensionType = "PHONE"
)

type ConversationAnalyticsGranularityType string

const (
	ConversationAnalyticsGranularityTypeHalfHour ConversationAnalyticsGranularityType = "HALF_HOUR"
	ConversationAnalyticsGranularityTypeDay      ConversationAnalyticsGranularityType = "DAILY"
	ConversationAnalyticsGranularityTypeMonth    ConversationAnalyticsGranularityType = "MONTHLY"
)

type ConversationAnalyticsOptions struct {
	Start        time.Time                            `json:"start" validate:"required"`
	End          time.Time                            `json:"end" validate:"required"`
	Granularity  ConversationAnalyticsGranularityType `json:"granularity" validate:"required"`
	PhoneNumbers []string                             `json:"phone_numbers,omitempty"`

	ConversationCategory  []ConversationCategoryType  `json:"conversation_category,omitempty"`
	ConversationTypes     []ConversationCategoryType  `json:"conversation_types,omitempty"`
	ConversationDirection []ConversationDirection     `json:"conversation_direction,omitempty"`
	Dimensions            []ConversationDimensionType `json:"dimensions,omitempty"`
}

type WhatsAppConversationAnalyticsNode struct {
	Start                 int    `json:"start" validate:"required"`
	End                   int    `json:"end,omitempty" validate:"required"`
	Conversation          int    `json:"conversation,omitempty"`
	PhoneNumber           string `json:"phone_number,omitempty"`
	Country               string `json:"country,omitempty"`
	ConversationType      string `json:"conversation_type,omitempty"`
	ConversationDirection string `json:"conversation_direction,omitempty"`
	ConversationCategory  string `json:"conversation_category,omitempty"`
	Cost                  int    `json:"cost,omitempty"`
}

type WhatsAppConversationAnalyticsEdge struct {
	Data []struct {
		DataPoints []WhatsAppConversationAnalyticsNode `json:"data_points,omitempty"`
	} `json:"data,omitempty"`
	Paging internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
}

type WhatsAppConversationAnalyticsResponse struct {
	ConversationAnalytics []WhatsAppConversationAnalyticsEdge `json:"conversation_analytics" validate:"required"`
}

// {"conversation_analytics":{"data":[{"data_points":[{"start":1701504000,"end":1701590400,"conversation":1,"cost":0},{"start":1701763200,"end":1701849600,"conversation":1,"cost":0},{"start":1715670000,"end":1715756400,"conversation":1,"cost":0},{"start":1716879600,"end":1716966000,"conversation":1,"cost":0},{"start":1716793200,"end":1716879600,"conversation":1,"cost":0},{"start":1713769200,"end":1713855600,"conversation":1,"cost":0},{"start":1713682800,"end":1713769200,"conversation":1,"cost":0},{"start":1716361200,"end":1716447600,"conversation":1,"cost":0},{"start":1715583600,"end":1715670000,"conversation":1,"cost":0},{"start":1717138800,"end":1717225200,"conversation":1,"cost":0},{"start":1716966000,"end":1717052400,"conversation":1,"cost":0},{"start":1717311600,"end":1717398000,"conversation":1,"cost":0},{"start":1713855600,"end":1713942000,"conversation":1,"cost":0},{"start":1716447600,"end":1716534000,"conversation":1,"cost":0}]}]},"id":"103043282674158"}

func (client *BusinessClient) ConversationAnalytics(options ConversationAnalyticsOptions) (*WhatsAppConversationAnalyticsResponse, error) {
	apiRequest := client.newBusinessApiRequest(client.BusinessAccountId, http.MethodGet)
	analyticsField := apiRequest.AddField(BusinessApiRequestQueryParamField{
		Name:    "conversation_analytics",
		Filters: map[string]string{},
	})
	analyticsField.AddFilter("start", fmt.Sprint(options.Start.Unix()))
	analyticsField.AddFilter("end", fmt.Sprint(options.End.Unix()))
	analyticsField.AddFilter("granularity", string(options.Granularity))

	if len(options.PhoneNumbers) > 0 {
		// get specific phone numbers
		analyticsField.AddFilter("phone_numbers", strings.Join(options.PhoneNumbers, ","))
	} else {
		// get all phone numbers
		analyticsField.AddFilter("phone_numbers", "[]")
	}

	if len(options.ConversationCategory) > 0 {
		categoryStrings := make([]string, len(options.ConversationCategory))
		for i, category := range options.ConversationCategory {
			categoryStrings[i] = string(category)
		}
		analyticsField.AddFilter("conversation_category", strings.Join(categoryStrings, ","))
	} else {
		analyticsField.AddFilter("conversation_category", "[]") // Empty slice
	}

	if len(options.ConversationTypes) > 0 {
		typeStrings := make([]string, len(options.ConversationTypes))
		for i, ctype := range options.ConversationTypes {
			typeStrings[i] = string(ctype)
		}
		analyticsField.AddFilter("conversation_types", strings.Join(typeStrings, ","))
	} else {
		analyticsField.AddFilter("conversation_types", "[]") // Empty slice
	}

	if len(options.ConversationDirection) > 0 {
		directionStrings := make([]string, len(options.ConversationDirection))
		for i, direction := range options.ConversationDirection {
			directionStrings[i] = string(direction)
		}
		analyticsField.AddFilter("conversation_direction", strings.Join(directionStrings, ","))
	} else {
		analyticsField.AddFilter("conversation_direction", "[]") // Empty slice
	}

	if len(options.Dimensions) > 0 {
		dimensionsStrings := make([]string, len(options.Dimensions))
		for i, dim := range options.Dimensions {
			dimensionsStrings[i] = string(dim)
		}
		analyticsField.AddFilter("dimensions", strings.Join(dimensionsStrings, ","))
	} else {
		// get all country codes
		analyticsField.AddFilter("dimensions", "[]")
	}

	response, err := apiRequest.Execute()
	if err != nil {
		// return wapi.go custom error here
		fmt.Println("Error while fetching business account", err)
	}
	var response_to_return WhatsAppConversationAnalyticsResponse
	json.Unmarshal([]byte(response), &response_to_return)

	fmt.Println("Response to return is", response_to_return)

	return &response_to_return, nil
}
