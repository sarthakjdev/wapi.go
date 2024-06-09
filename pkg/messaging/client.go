package messaging

import "github.com/sarthakjdev/wapi.go/internal/manager"

// MessagingClient represents a WhatsApp client.
type MessagingClient struct {
	Media             manager.MediaManager
	Message           manager.MessageManager
	PhoneNumberId     string
	ApiAccessToken    string
	BusinessAccountId string
}

// GetPhoneNumberId returns the phone number ID associated with the client.
func (client *MessagingClient) GetPhoneNumberId() string {
	return client.PhoneNumberId
}

// SetPhoneNumberId sets the phone number ID for the client.
func (client *MessagingClient) SetPhoneNumberId(phoneNumberId string) {
	client.PhoneNumberId = phoneNumberId
}

func (client *MessagingClient) GetApiAccessToken() string {
	return client.ApiAccessToken
}

func (client *MessagingClient) SetApiAccessToken(apiAccessToken string) {
	client.ApiAccessToken = apiAccessToken
}

func (client *MessagingClient) GetBusinessAccountId() string {
	return client.BusinessAccountId
}
