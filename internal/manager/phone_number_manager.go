package manager

import "github.com/sarthakjdev/wapi.go/internal"

type PhoneNumber struct {
}

type PhoneNumberManager struct {
}

type PhoneNumberManagerConfig struct {
}

func NewPhoneNumberManager(config *PhoneNumberManagerConfig) *PhoneNumberManager {
	return &PhoneNumberManager{}
}

type WhatsappBusinessAccountPhoneNumber struct {
}

type WhatsappBusinessAccountPhoneNumberEdge struct {
	Data    []WhatsappBusinessAccountPhoneNumber       `json:"data,omitempty"`
	Paging  internal.WhatsAppBusinessApiPaginationMeta `json:"paging,omitempty"`
	Summary string                                     `json:"summary,omitempty"`
}

func (pm *PhoneNumberManager) FetchAll() {
	// ! TODO: call this API endpoint
	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account/phone_numbers/
}

// Fetch fetches a phone number by its ID.
func (pm *PhoneNumberManager) Fetch(Id string) {

	// ! TODO: call this API endpoint

	// https: //developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/

}

func (pm *PhoneNumberManager) Add() {
	// ! TODO: call this API endpoint to create a new phone number for the whats app business accounts

	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/#:~:text=Graph%20API%20Version,the%20following%20paths%3A
}

func (pm *PhoneNumberManager) Update() {
	// ! TODO: call this APO endpoint to update the phone number for the whats app business accounts

	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/#:~:text=Edit%20failure-,Updating,-You%20can%20update
}

func (pm *PhoneNumberManager) UpdateCartAndCatalogSettings() {

	// ! TODO: call this API endpoint to update the cart and catalog settings for the whats app business accounts phone number

	// https://developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/whatsapp_commerce_settings
}

func (pm *PhoneNumberManager) RequestCodeForVerification() {
	// https: //developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/request_code/
}

func (pm *PhoneNumberManager) VerifyCode() {
	// https: //developers.facebook.com/docs/graph-api/reference/whats-app-business-account-to-number-current-status/verify_code
}
