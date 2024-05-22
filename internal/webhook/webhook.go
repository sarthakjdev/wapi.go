package webhook

import (
	"github.com/labstack/echo/v4"
)

// references for event driven architecture in golang:  https://medium.com/@souravchoudhary0306/implementation-of-event-driven-architecture-in-go-golang-28d9a1c01f91
type Webhook struct {
	secret string
	path   string
	port   int
}

type WebhookManagerConfig struct {
	Secret string
	Path   string
	Port   int
}

func NewWebhook(options WebhookManagerConfig) *Webhook {
	return &Webhook{
		secret: options.Secret,
		path:   options.Path,
		port:   options.Port,
	}
}

// this function is used in case if the client have not provided any custom http server
func (wh *Webhook) createEchoHttpServer() *echo.Echo {
	e := echo.New()
	return e

}

func (wh *Webhook) ListenToEvents() {

}
