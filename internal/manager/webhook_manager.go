package manager

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

// references for event driven architecture in golang:  https://medium.com/@souravchoudhary0306/implementation-of-event-driven-architecture-in-go-golang-28d9a1c01f91
type WebhookManager struct {
	secret       string
	path         string
	port         int
	EventManager EventManger
	Requester    RequestClient
}

type WebhookManagerConfig struct {
	Secret       string
	Path         string
	Port         int
	EventManager EventManger
	Requester    RequestClient
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

func (wh *WebhookManager) getRequestHandler(req *http.Request) {
}

func (wh *WebhookManager) postRequestHandler(req *http.Request) {
	// emits events based on the payload of the request

	wh.EventManager.Publish(TextMessageEvent, events.NewTextMessageEvent(
		"wiuhbiueqwdqwd",
		"2134141414",
		"hello",
	))

}

func (wh *WebhookManager) ListenToEvents() {

	fmt.Println("Listening to events")
	server := wh.createEchoHttpServer()

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
