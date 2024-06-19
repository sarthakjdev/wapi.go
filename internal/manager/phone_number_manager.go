package manager

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sarthakjdev/wapi.go/internal"
	"github.com/sarthakjdev/wapi.go/internal/request_client"
)

type PhoneNumberManager struct {
	businessAccountId string
	aApiAccessToken   string
	requester         *request_client.RequestClient
}

type PhoneNumberManagerConfig struct {
	BusinessAccountId string
	ApiAccessToken    string
	Requester         *request_client.RequestClient
}

func NewPhoneNumberManager(config *PhoneNumberManagerConfig) *PhoneNumberManager {
	return &PhoneNumberManager{
		aApiAccessToken:   config.ApiAccessToken,
		businessAccountId: config.BusinessAccountId,
		requester:         config.Requester,
	}
}

type WhatsappBusinessAccountPhoneNumber struct {
	VerifiedName       string `json:"verified_name,omitempty"`
	DisplayPhoneNumber string `json:"display_phone_number,omitempty"`
	Id                 string `json:"id,omitempty"`
	QualityRating      string `json:"quality_rating,omitempty"`
	CodeVerification   struct {
		Status string `json:"code_verification_status,omitempty"`
	} `json:"code_verification_status,omitempty"`
	PlatformType string `json:"platform_type,omitempty"`
}

type WhatsappBusinessAccountPhoneNumberEdge struct {
	Data    []WhatsappBusinessAccountPhoneNumber       `json:"data,omitempty"`
	Paging  internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
	Summary string                                     `json:"summary,omitempty"`
}

type FetchPhoneNumberFilters struct {
	GetSandboxNumbers bool
}

func (manager *PhoneNumberManager) FetchAll(options FetchPhoneNumberFilters) (*WhatsappBusinessAccountPhoneNumberEdge, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{manager.businessAccountId, "/", "phone_numbers"}, ""), http.MethodGet)

	apiRequest.AddQueryParam("filtering", `[{"field":"account_mode","operator":"EQUAL","value":"LIVE"}]`)
	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}

	var response_to_return WhatsappBusinessAccountPhoneNumberEdge
	json.Unmarshal([]byte(response), &response_to_return)

	return &response_to_return, nil
}

// Fetch fetches a phone number by its ID.
func (manager *PhoneNumberManager) Fetch(phoneNumberId string) (*WhatsappBusinessAccountPhoneNumber, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(phoneNumberId, http.MethodGet)

	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}

	var response_to_return WhatsappBusinessAccountPhoneNumber
	json.Unmarshal([]byte(response), &response_to_return)

	return &response_to_return, nil

}

type GenerateQrCodeResponse struct {
	Code             string `json:"code,omitempty"`
	PrefilledMessage string `json:"prefilled_message,omitempty"`
	DeepLinkUrl      string `json:"deep_link_url,omitempty"`
	QrImageUrl       string `json:"qr_image_url,omitempty"`
}

func (manager *PhoneNumberManager) GenerateQrCode(phoneNumber string, prefilledMessage string) (*GenerateQrCodeResponse, error) {

	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{phoneNumber, "/message_qrdls"}, ""), http.MethodPost)
	jsonBody, err := json.Marshal(map[string]string{
		"prefilled_message": prefilledMessage,
		"generate_qr_image": "PNG",
	})
	if err != nil {
		return nil, err
	}
	apiRequest.SetBody(string(jsonBody))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var response_to_return GenerateQrCodeResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

type GetAllQrCodesResponse struct {
	Data []GenerateQrCodeResponse `json:"data,omitempty"`
}

func (manager *PhoneNumberManager) GetAllQrCodes(phoneNumber string) (*GetAllQrCodesResponse, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{phoneNumber, "/message_qrdls"}, ""), http.MethodGet)
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var response_to_return GetAllQrCodesResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

func (manager *TemplateManager) GetQrCodeById(phoneNumber, id string) (*GetAllQrCodesResponse, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{phoneNumber, "/message_qrdls", "/", id}, ""), http.MethodDelete)
	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}
	var response_to_return GetAllQrCodesResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

type DeleteQrCodeResponse struct {
	Success bool `json:"success,omitempty"`
}

func (manager *PhoneNumberManager) DeleteQrCode(phoneNumber, id string) (*DeleteQrCodeResponse, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{phoneNumber, "/message_qrdls", "/", id}, ""), http.MethodDelete)
	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}
	var response_to_return DeleteQrCodeResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

func (manager *PhoneNumberManager) UpdateQrCode(phoneNumber, id, prefilledMessage string) (*GenerateQrCodeResponse, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{phoneNumber, "/message_qrdls"}, ""), http.MethodPost)
	jsonBody, err := json.Marshal(map[string]string{
		"prefilled_message": prefilledMessage,
		"code":              id,
	})
	if err != nil {
		return nil, err
	}
	apiRequest.SetBody(string(jsonBody))
	response, err := apiRequest.Execute()
	if err != nil {
		return nil, err
	}

	var response_to_return GenerateQrCodeResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}
