// Package manager provides functionality for managing webhooks.
package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthakjdev/wapi.go/internal"
	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

// WebhookManager represents a manager for handling webhooks.
type WebhookManager struct {
	secret       string
	path         string
	port         int
	EventManager EventManger
	Requester    requestclient.RequestClient
}

// WebhookManagerConfig represents the configuration options for creating a new WebhookManager.
type WebhookManagerConfig struct {
	Secret       string
	Path         string
	Port         int
	EventManager EventManger
	Requester    requestclient.RequestClient
}

// NewWebhook creates a new WebhookManager with the given options.
func NewWebhook(options *WebhookManagerConfig) *WebhookManager {
	return &WebhookManager{
		secret:       options.Secret,
		path:         options.Path,
		port:         options.Port,
		EventManager: options.EventManager,
		Requester:    options.Requester,
	}
}

// createEchoHttpServer creates a new instance of Echo HTTP server.
// This function is used in case the client has not provided any custom HTTP server.
func (wh *WebhookManager) createEchoHttpServer() *echo.Echo {
	e := echo.New()
	return e
}

// GetRequestHandler handles GET requests to the webhook endpoint.
func (wh *WebhookManager) GetRequestHandler(c echo.Context) error {
	hubVerificationToken := c.QueryParam("hub.verify_token")
	hubChallenge := c.QueryParam("hub.challenge")
	fmt.Println(hubVerificationToken, hubChallenge)
	if hubVerificationToken == wh.secret {
		return c.String(200, hubChallenge)
	} else {
		return c.String(400, "invalid token")
	}
}

// PostRequestHandler handles POST requests to the webhook endpoint.
func (wh *WebhookManager) PostRequestHandler(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		c.String(400, "error reading request body")
	}

	var payload WhatsappApiNotificationPayloadSchemaType

	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		c.String(400, "Invalid JSON data")
	}

	if err := internal.GetValidator().Struct(payload); err != nil {
		fmt.Println("Error validating JSON:", err)
		c.String(400, "Invalid JSON data")
	}

	for _, entry := range payload.Entry {
		for _, change := range entry.Changes {
			for _, message := range change.Value.Messages {
				switch message.Type {
				case NotificationMessageTypeText:
					{
						wh.EventManager.Publish(events.TextMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeImage:
					{

						wh.EventManager.Publish(events.ImageMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)

					}
				case NotificationMessageTypeVideo:
					{

						wh.EventManager.Publish(events.AudioMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)

					}
				case NotificationMessageTypeDocument:
					{

						wh.EventManager.Publish(events.DocumentMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeAudio:
					{
						wh.EventManager.Publish(events.AudioMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeLocation:
					{
						wh.EventManager.Publish(events.LocationMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeContacts:
					{
						wh.EventManager.Publish(events.ContactMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeSticker:
					{
						wh.EventManager.Publish(events.StickerMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeSystem:
					{
						wh.EventManager.Publish(events.TextMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeButton:
					{
						wh.EventManager.Publish(events.ReplyButtonInteractionEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeInteractive:
					{
						wh.EventManager.Publish(events.ListInteractionMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeReaction:
					{
						wh.EventManager.Publish(events.ReactionMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeOrder:
					{
						wh.EventManager.Publish(events.OrderReceivedEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeUnknown:
					{
						wh.EventManager.Publish(events.UnknownEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								message.Id,
								"",
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				}

			}
		}
	}

	c.String(200, "Message received")

	fmt.Println("Received valid payload:", payload.Entry[0].Changes[0].Value.Messages[0].Type)

	return nil
}

// ListenToEvents starts listening to events and handles incoming requests.
func (wh *WebhookManager) ListenToEvents() {
	fmt.Println("Listening to events")
	server := wh.createEchoHttpServer()
	server.GET(wh.path, wh.GetRequestHandler)
	server.POST(wh.path, wh.PostRequestHandler)

	// Start server in a goroutine
	go func() {
		if err := server.Start(":8080"); err != nil {
			return
		}
	}()

	wh.EventManager.Publish(events.ReadyEventType, events.NewReadyEvent())
	// Wait for an interrupt signal (e.g., Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // Capture SIGINT (Ctrl+C)
	<-quit                            // Wait for the signal

	// Gracefully shut down the server (optional)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err) // Handle shutdown errors gracefully
	}
}
