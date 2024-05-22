package wapi

import (
	"fmt"

	manager "github.com/sarthakjdev/wapi.go/internal/manager"
	"github.com/sarthakjdev/wapi.go/internal/webhook"
	"github.com/sarthakjdev/wapi.go/utils"
)

// Client represents a WhatsApp client.
type Client struct {
	Media             manager.MediaManager
	Message           manager.MessageManager
	Phone             manager.PhoneNumbersManager
	webhook           webhook.Webhook
	phoneNumberId     string
	apiAccessToken    string
	businessAccountId string
}

// ClientConfig represents the configuration options for the WhatsApp client.
type ClientConfig struct {
	PhoneNumberId     string `validate:"required"`
	ApiAccessToken    string `validate:"required"`
	BusinessAccountId string `validate:"required"`
	WebhookPath       string `validate:"required"`
	WebhookSecret     string `validate:"required"`
	WebhookServerPort int
}

// NewWapiClient creates a new instance of Client.
func New(configs ClientConfig) *Client {
	err := utils.GetValidator().Struct(configs)
	if err != nil {
		fmt.Println("error validating client config", err)
		return nil
	}
	requester := *manager.NewRequestClient(configs.PhoneNumberId, configs.ApiAccessToken)
	return &Client{
		Media:             *manager.NewMediaManager(requester),
		Message:           *manager.NewMessageManager(requester),
		Phone:             *manager.NewPhoneNumbersManager(requester),
		webhook:           *webhook.NewWebhook(webhook.WebhookManagerConfig{Path: configs.WebhookPath, Secret: configs.WebhookSecret, Port: configs.WebhookServerPort}),
		phoneNumberId:     configs.PhoneNumberId,
		apiAccessToken:    configs.ApiAccessToken,
		businessAccountId: configs.BusinessAccountId,
	}
}

// GetPhoneNumberId returns the phone number ID associated with the client.
func (client *Client) GetPhoneNumberId() string {
	return client.phoneNumberId
}

// SetPhoneNumberId sets the phone number ID for the client.
func (client *Client) SetPhoneNumberId(phoneNumberId string) {
	client.phoneNumberId = phoneNumberId
}

// InitiateClient initializes the client and starts listening to events from the webhook.
// It returns true if the client was successfully initiated.
func (client *Client) InitiateClient() bool {
	client.webhook.ListenToEvents()
	return true
}
