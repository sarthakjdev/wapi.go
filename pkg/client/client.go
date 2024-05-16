package wapi

import (
	"github.com/sarthakjdev/wapi.go/pkg/webhook"
)

// Client represents a WhatsApp client.
type Client struct {
	Media             MediaManager
	message           MessageManager
	Phone             PhoneNumbersManager
	webhook           webhook.Webhook
	phoneNumberId     string
	apiAccessToken    string
	host              string
	businessAccountId string
	apiVersion        string
}

type ClientConfig struct {
	phoneNumberId     string
	apiAccessToken    string
	businessAccountId string
	webhookPath       string
	webhookSecret     string
	webhookServerPort int
}

// NewClient creates a new instance of Client.
func NewWapiClient(options ClientConfig) *Client {
	requester := *NewRequestClient()

	return &Client{
		Media:             *NewMediaManager(requester),
		message:           *NewMessageManager(requester),
		Phone:             *NewPhoneNumbersManager(requester),
		webhook:           *webhook.NewWebhook(webhook.WebhookManagerConfig{Path: options.webhookPath, Secret: options.webhookSecret, Port: options.webhookServerPort}),
		phoneNumberId:     options.phoneNumberId,
		apiAccessToken:    options.apiAccessToken,
		host:              "graph.facebook.com",
		businessAccountId: options.businessAccountId,
		apiVersion:        "v19.0",
	}
}

func (client *Client) GetPhoneNumberId() string {
	return client.phoneNumberId
}

func (client *Client) SetPhoneNumberId(phoneNumberId string) {
	client.phoneNumberId = phoneNumberId
}

func (client *Client) InitiateClient() bool {
	client.webhook.ListenToEvents()
	return true
}
