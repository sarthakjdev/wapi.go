package messaging

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sarthakjdev/wapi.go/internal"
	"github.com/sarthakjdev/wapi.go/internal/manager"
	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

// Client represents a WhatsApp client.
type Client struct {
	Media             manager.MediaManager
	Message           manager.MessageManager
	webhook           manager.WebhookManager
	phoneNumberId     string
	apiAccessToken    string
	businessAccountId string
}

// ClientConfig represents the configuration options for the WhatsApp client.
type ClientConfig struct {
	PhoneNumberId     string `validate:"required"`
	ApiAccessToken    string `validate:"required"`
	BusinessAccountId string `validate:"required"`
	WebhookSecret     string `validate:"required"`

	// these two are not required, because may be user want to use their own server
	WebhookPath       string
	WebhookServerPort int
}

// NewWapiClient creates a new instance of Client.
func New(configs ClientConfig) (*Client, error) {
	// Validate the client configuration options
	err := internal.GetValidator().Struct(configs)
	if err != nil {
		return nil, fmt.Errorf("error validating client config: %w", err)
	}

	// Create a new request client
	requester := *requestclient.NewRequestClient(configs.PhoneNumberId, configs.ApiAccessToken)

	// Create a new event manager
	eventManager := *manager.NewEventManager()

	// Create a new Client instance with the provided configurations
	return &Client{
		Media:             *manager.NewMediaManager(requester),
		Message:           *manager.NewMessageManager(requester),
		webhook:           *manager.NewWebhook(&manager.WebhookManagerConfig{Path: configs.WebhookPath, Secret: configs.WebhookSecret, Port: configs.WebhookServerPort, EventManager: eventManager, Requester: requester}),
		phoneNumberId:     configs.PhoneNumberId,
		apiAccessToken:    configs.ApiAccessToken,
		businessAccountId: configs.BusinessAccountId,
	}, nil
}

// GetPhoneNumberId returns the phone number ID associated with the client.
func (client *Client) GetPhoneNumberId() string {
	return client.phoneNumberId
}

// SetPhoneNumberId sets the phone number ID for the client.
func (client *Client) SetPhoneNucmberId(phoneNumberId string) {
	client.phoneNumberId = phoneNumberId
}

// InitiateClient initializes the client and starts listening to events from the webhook.
// It returns true if the client was successfully initiated.
func (client *Client) InitiateClient() bool {
	client.webhook.ListenToEvents()
	return true
}

// GetWebhookGetRequestHandler returns the handler function for handling GET requests to the webhook.
func (client *Client) GetWebhookGetRequestHandler() func(c echo.Context) error {
	return client.webhook.GetRequestHandler
}

// GetWebhookPostRequestHandler returns the handler function for handling POST requests to the webhook.
func (client *Client) GetWebhookPostRequestHandler() func(c echo.Context) error {
	return client.webhook.PostRequestHandler
}

// OnMessage registers a handler for a specific event type.
func (client *Client) On(eventType events.EventType, handler func(events.BaseEvent)) {
	client.webhook.EventManager.On(eventType, handler)
}
