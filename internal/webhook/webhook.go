package webhook

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

func (wh *Webhook) ListenToEvents() {

}
