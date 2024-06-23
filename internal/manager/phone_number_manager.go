// Package manager provides functionality to manage phone numbers for WhatsApp Business API.
package manager

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sarthakjdev/wapi.go/internal"
	"github.com/sarthakjdev/wapi.go/internal/request_client"
)

// PhoneNumberManager is responsible for managing phone numbers for WhatsApp Business API and phone number specific operations.
type PhoneNumberManager struct {
	businessAccountId string
	aApiAccessToken   string
	requester         *request_client.RequestClient
}

// PhoneNumberManagerConfig holds the configuration for PhoneNumberManager.
type PhoneNumberManagerConfig struct {
	BusinessAccountId string
	ApiAccessToken    string
	Requester         *request_client.RequestClient
}

// NewPhoneNumberManager creates a new instance of PhoneNumberManager.
func NewPhoneNumberManager(config *PhoneNumberManagerConfig) *PhoneNumberManager {
	return &PhoneNumberManager{
		aApiAccessToken:   config.ApiAccessToken,
		businessAccountId: config.BusinessAccountId,
		requester:         config.Requester,
	}
}

// WhatsappBusinessAccountPhoneNumber represents a WhatsApp Business Account phone number.
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

// WhatsappBusinessAccountPhoneNumberEdge represents a list of WhatsApp Business Account phone numbers.
type WhatsappBusinessAccountPhoneNumberEdge struct {
	Data    []WhatsappBusinessAccountPhoneNumber       `json:"data,omitempty"`
	Paging  internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
	Summary string                                     `json:"summary,omitempty"`
}

// FetchPhoneNumberFilters holds the filters for fetching phone numbers.
type FetchPhoneNumberFilters struct {
	GetSandboxNumbers bool
}

// FetchAll fetches all phone numbers based on the provided filters.
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

// GenerateQrCodeResponse represents the response of generating a QR code.
type GenerateQrCodeResponse struct {
	Code             string `json:"code,omitempty"`
	PrefilledMessage string `json:"prefilled_message,omitempty"`
	DeepLinkUrl      string `json:"deep_link_url,omitempty"`
	QrImageUrl       string `json:"qr_image_url,omitempty"`
}

// GenerateQrCode generates a QR code for the specified phone number with the given prefilled message.
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

// GetAllQrCodesResponse represents the response of getting all QR codes for a phone number.
type GetAllQrCodesResponse struct {
	Data []GenerateQrCodeResponse `json:"data,omitempty"`
}

// GetAllQrCodes gets all QR codes for the specified phone number.
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

// GetQrCodeById gets a QR code by its ID for the specified phone number.
func (manager *PhoneNumberManager) GetQrCodeById(phoneNumber, id string) (*GetAllQrCodesResponse, error) {
	apiRequest := manager.requester.NewBusinessApiRequest(strings.Join([]string{phoneNumber, "/message_qrdls", "/", id}, ""), http.MethodDelete)
	response, err := apiRequest.Execute()

	if err != nil {
		return nil, err
	}
	var response_to_return GetAllQrCodesResponse
	json.Unmarshal([]byte(response), &response_to_return)
	return &response_to_return, nil
}

// DeleteQrCodeResponse represents the response of deleting a QR code.
type DeleteQrCodeResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteQrCode deletes a QR code by its ID for the specified phone number.
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

// UpdateQrCode updates a QR code by its ID for the specified phone number with the given prefilled message.
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
