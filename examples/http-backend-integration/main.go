package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sarthakjdev/wapi.go/internal/manager"
	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
	"github.com/sarthakjdev/wapi.go/pkg/components"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

func main() {

	whatsappClient, err := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "",
		ApiAccessToken:    "",
		BusinessAccountId: "",
		WebhookPath:       "/webhook",
		WebhookSecret:     "1234567890",
		WebhookServerPort: 8080,
	})

	whatsappClient.On(manager.TextMessageEvent, func(event events.BaseEvent) {
		textMessageEvent := event.(*events.TextMessageEvent)
		reply, err := components.NewTextMessage(components.TextMessageConfigs{
			Text: "Hello, from wapi.go",
		})
		if err != nil {
			fmt.Println("error creating text message", err)
			return
		}
		textMessageEvent.Reply(reply)
	})

	getHandler := whatsappClient.GetWebhookGetRequestHandler()
	postHandler := whatsappClient.GetWebhookPostRequestHandler()

	if err != nil {
		fmt.Println("error creating client", err)
		return
	}

	server := echo.New()

	server.GET("/webhook", getHandler)
	server.POST("/webhook", postHandler)

	server.Start(":8080")

}
