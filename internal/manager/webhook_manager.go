package manager

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

// references for event driven architecture in golang:  https://medium.com/@souravchoudhary0306/implementation-of-event-driven-architecture-in-go-golang-28d9a1c01f91
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

func (wh *WebhookManager) getRequestHandler(c echo.Context) {

	// this endpoint is used to verify the webhook

	request := c.Request()
	fmt.Println(request)

}

func (wh *WebhookManager) postRequestHandler(c echo.Context) {
	// emits events based on the payload of the request

	request := c.Request()

	fmt.Println(request)

	// parse the request here
	// get the type of message
	// emit the event based on the type of message

	wh.EventManager.Publish(TextMessageEvent, events.NewTextMessageEvent(
		"wiuhbiueqwdqwd",
		"2134141414",
		"hello",
		wh.Requester,
	))

}

func (wh *WebhookManager) ListenToEvents() {

	fmt.Println("Listening to events")
	server := wh.createEchoHttpServer()

	server.GET(wh.path, func(c echo.Context) error {
		wh.getRequestHandler(c)
		return c.String(200, "ok")
	})

	server.POST(wh.path, func(c echo.Context) error {
		wh.postRequestHandler(c)
		return c.String(200, "ok")
	})

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
