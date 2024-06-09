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
	"github.com/sarthakjdev/wapi.go/pkg/components"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

// WebhookManager represents a manager for handling webhooks.
type WebhookManager struct {
	secret       string
	path         string
	port         int
	EventManager EventManager
	Requester    requestclient.RequestClient
}

// WebhookManagerConfig represents the configuration options for creating a new WebhookManager.
type WebhookManagerConfig struct {
	Secret       string                      `validate:"required"`
	EventManager EventManager                `validate:"required"`
	Requester    requestclient.RequestClient `validate:"required"`
	Path         string
	Port         int
}

// NewWebhook creates a new WebhookManager with the given options.
func NewWebhook(options *WebhookManagerConfig) *WebhookManager {
	if err := internal.GetValidator().Struct(options); err != nil {
		fmt.Println("Error validating WebhookManagerConfig:", err)
		return nil
	}
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

			if len(change.Value.Statuses) > 0 {
				for _, status := range change.Value.Statuses {

					switch status.Status {
					case string(MessageStatusDelivered):
						{
							wh.EventManager.Publish(events.MessageDeliveredEventType, events.NewMessageDeliveredEvent(events.BaseSystemEvent{
								Timestamp: status.Timestamp,
							}, status.Conversation.Id, status.RecipientId))
						}

					case string(MessageStatusRead):
						{
							wh.EventManager.Publish(events.MessageReadEventType, events.NewMessageReadEvent(events.BaseSystemEvent{
								Timestamp: status.Timestamp,
							}, status.Conversation.Id, status.RecipientId))
						}
					case string(MessageStatusSent):
						{
							wh.EventManager.Publish(events.MessageSentEventType, events.NewMessageSentEvent(events.BaseSystemEvent{
								Timestamp: status.Timestamp,
							}, status.Conversation.Id, status.RecipientId))
						}
					case string(MessageStatusFailed):
						{
							// ! TODO: check and properly emit the error event here.
						}
					}

				}
			}

			// this is the id of the phone number which can ve used to respond to the events received
			phoneNumberId := change.Value.Metadata.PhoneNumberId

			for _, message := range change.Value.Messages {
				switch message.Type {
				case NotificationMessageTypeText:
					{
						wh.EventManager.Publish(events.TextMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeImage:
					{

						imageMessageComponent, err := components.NewImageMessage(components.ImageMessageConfigs{
							Id:      message.Image.Id,
							Caption: message.Image.Caption,
						})

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating image message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.ImageMessageEventType, events.NewImageMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*imageMessageComponent,
							message.Image.MIMEType, message.Image.SHA256, message.Image.Id),
						)
					}
				case NotificationMessageTypeAudio:
					{

						audioMessageComponent, err := components.NewAudioMessage(components.AudioMessageConfigs{
							Id: message.Audio.Id,
						})

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating audio message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.AudioMessageEventType, events.NewAudioMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*audioMessageComponent,
							message.Audio.MIMEType, message.Audio.SHA256, message.Audio.Id),
						)

					}
				case NotificationMessageTypeVideo:
					{

						videoMessageComponent, err := components.NewVideoMessage(components.VideoMessageConfigs{
							Id:      message.Video.Id,
							Caption: message.Video.Caption,
						})

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating Video message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.VideoMessageEventType, events.NewVideoMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*videoMessageComponent,
							message.Video.MIMEType, message.Video.SHA256, message.Video.Id),
						)

					}
				case NotificationMessageTypeDocument:
					{
						documentMessageComponent, err := components.NewVideoMessage(components.VideoMessageConfigs{
							Id:      message.Document.Id,
							Caption: message.Document.Caption,
						})

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating document message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.DocumentMessageEventType, events.NewVideoMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*documentMessageComponent,
							message.Document.MIMEType, message.Document.SHA256, message.Document.Id),
						)
					}
				case NotificationMessageTypeLocation:
					{
						locationMessageComponent, err := components.NewLocationMessage(message.Location.Latitude, message.Location.Longitude)

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating location message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.LocationMessageEventType, events.NewLocationMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*locationMessageComponent),
						)
					}
				case NotificationMessageTypeContacts:
					{
						wh.EventManager.Publish(events.ContactMessageEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeSticker:
					{

						stickerMessageComponent, err := components.NewStickerMessage(&components.StickerMessageConfigs{
							Id: message.Sticker.Id,
						})

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating Sticker message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.StickerMessageEventType, events.NewStickerMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*stickerMessageComponent,
							message.Sticker.MIMEType, message.Sticker.SHA256, message.Sticker.Id),
						)

					}
				case NotificationMessageTypeButton:
					{
						wh.EventManager.Publish(events.QuickReplyMessageEventType, events.NewQuickReplyButtonInteractionEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Button.Text,
							message.Button.Payload,
						))
					}
				case NotificationMessageTypeInteractive:
					{
						if message.Interactive.Type == "list" {
							wh.EventManager.Publish(events.ListInteractionMessageEventType, events.NewListInteractionEvent(
								events.NewBaseMessageEvent(
									phoneNumberId,
									message.Id,
									message.Timestamp,
									message.From,
									message.Context.Forwarded,
									wh.Requester),
								message.Interactive.ListReply.Title,
								message.Interactive.ListReply.Id,
								message.Interactive.ListReply.Description,
							))
						} else {
							wh.EventManager.Publish(events.ReplyButtonInteractionEventType, events.NewReplyButtonInteractionEvent(
								events.NewBaseMessageEvent(
									phoneNumberId,
									message.Id,
									message.Timestamp,
									message.From,
									message.Context.Forwarded,
									wh.Requester),
								message.Interactive.ButtonReply.Title,
								message.Interactive.ButtonReply.ReplyId,
							))
						}
					}
				case NotificationMessageTypeReaction:
					{
						reactionMessageComponent, err := components.NewReactionMessage(components.ReactionMessageParams{
							MessageId: message.Reaction.MessageId,
							Emoji:     message.Reaction.Emoji,
						})

						if err != nil {
							// ! TODO: emit error event here
							fmt.Println("Error creating location message:", err)
							c.String(500, "Internal server error")
							return err
						}

						wh.EventManager.Publish(events.ReactionMessageEventType, events.NewReactionMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							*reactionMessageComponent,
						))
					}
				case NotificationMessageTypeOrder:
					{
						wh.EventManager.Publish(events.OrderReceivedEventType, events.NewTextMessageEvent(
							events.NewBaseMessageEvent(
								phoneNumberId,
								message.Id,
								message.Timestamp,
								message.From,
								message.Context.Forwarded,
								wh.Requester),
							message.Text.Body),
						)
					}
				case NotificationMessageTypeSystem:
					{
						if message.System.Type == SystemNotificationTypeCustomerIdentityChanged {
							wh.EventManager.Publish(events.CustomerIdentityChangedEventType, events.CustomerIdentityChangedEvent{
								BaseSystemEvent: events.BaseSystemEvent{
									Timestamp: message.Timestamp,
								},
								Acknowledged:      message.Identity.Acknowledged,
								CreationTimestamp: message.Identity.CreatedTimestamp,
								Hash:              message.Identity.Hash,
							})
						} else {
							wh.EventManager.Publish(events.CustomerNumberChangedEventType, events.CustomerNumberChangedEvent{
								BaseSystemEvent: events.BaseSystemEvent{
									Timestamp: message.Timestamp,
								},
								NewWaId:           message.System.WaId,
								OldWaId:           message.System.Customer,
								ChangeDescription: message.System.Body,
							})
						}
					}
				case NotificationMessageTypeUnknown:
					{
						// ! TODO: handle error in the event and then emit it.
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
