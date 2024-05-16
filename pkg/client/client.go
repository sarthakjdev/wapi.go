package wapi

import (
	manager "github.com/sarthakjdev/wapi.go/internal/manager"
	"github.com/sarthakjdev/wapi.go/internal/webhook"
)

// Client represents a WhatsApp client.
type Client struct {
	Media             manager.MediaManager
	message           manager.MessageManager
	Phone             manager.PhoneNumbersManager
	webhook           webhook.Webhook
	phoneNumberId     string
	apiAccessToken    string
	host              string
	businessAccountId string
	apiVersion        string
}

// ClientConfig represents the configuration options for the WhatsApp client.
type ClientConfig struct {
	phoneNumberId     string
	apiAccessToken    string
	businessAccountId string
	webhookPath       string
	webhookSecret     string
	webhookServerPort int
}

// NewWapiClient creates a new instance of Client.
func NewWapiClient(options ClientConfig) *Client {
	requester := *manager.NewRequestClient()

	return &Client{
		Media:             *manager.NewMediaManager(requester),
		message:           *manager.NewMessageManager(requester),
		Phone:             *manager.NewPhoneNumbersManager(requester),
		webhook:           *webhook.NewWebhook(webhook.WebhookManagerConfig{Path: options.webhookPath, Secret: options.webhookSecret, Port: options.webhookServerPort}),
		phoneNumberId:     options.phoneNumberId,
		apiAccessToken:    options.apiAccessToken,
		host:              "graph.facebook.com",
		businessAccountId: options.businessAccountId,
		apiVersion:        "v19.0",
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
