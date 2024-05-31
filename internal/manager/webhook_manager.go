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
	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/events"
	"github.com/sarthakjdev/wapi.go/utils"
)

type WebhookManager struct {
	secret       string
	path         string
	port         int
	EventManager EventManger
	Requester    requestclient.RequestClient
}

type WebhookManagerConfig struct {
	Secret       string
	Path         string
	Port         int
	EventManager EventManger
	Requester    requestclient.RequestClient
}

func NewWebhook(options *WebhookManagerConfig) *WebhookManager {
	return &WebhookManager{
		secret:       options.Secret,
		path:         options.Path,
		port:         options.Port,
		EventManager: options.EventManager,
		Requester:    options.Requester,
	}
}

// this function is used in case if the client have not provided any custom http server
func (wh *WebhookManager) createEchoHttpServer() *echo.Echo {
	e := echo.New()
	return e

}

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

	if err := utils.GetValidator().Struct(payload); err != nil {
		fmt.Println("Error validating JSON:", err)
		c.String(400, "Invalid JSON data")
	}

	for _, entry := range payload.Entry {
		for _, change := range entry.Changes {
			for _, message := range change.Value.Messages {
				switch message.Type {
				case NotificationMessageTypeText:
					{
						wh.EventManager.Publish(TextMessageEvent, events.NewReactionMessageEvent(
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

						wh.EventManager.Publish(ImageMessageEvent, events.NewReactionMessageEvent(
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

						wh.EventManager.Publish(AudioMessageEvent, events.NewReactionMessageEvent(
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

						wh.EventManager.Publish(DocumentMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(AudioMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(LocationMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(ContactMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(StickerMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(TextMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(ReplyButtonInteractionEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(ListInteractionMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(ReactionMessageEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(OrderReceivedEvent, events.NewReactionMessageEvent(
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
						wh.EventManager.Publish(UnknownEvent, events.NewReactionMessageEvent(
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

	wh.EventManager.Publish(ReadyEvent, events.NewReadyEvent())
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

func (wh *WebhookManager) determineEventType() EventType {
	return TextMessageEvent
}
