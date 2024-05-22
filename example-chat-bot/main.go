package main

import (
	"github.com/sarthakjdev/wapi.go/internal/manager"
	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
	wapiModels "github.com/sarthakjdev/wapi.go/pkg/models"
)

func main() {

	// ! TODO: programmatic read the env variables here

	// creating a client
	whatsappClient, err := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "113269274970227",
		ApiAccessToken:    "",
		BusinessAccountId: "103043282674158",
		WebhookPath:       "/webhook",
		WebhookSecret:     "123456789",
		WebhookServerPort: 8080,
	})

	if err != nil {
		return
	}

	// create a message
	textMessage, err := wapiModels.NewTextMessage(wapiModels.TextMessageConfigs{
		Text:         "Hello, from wapi.go",
		AllowPreview: true,
	})

	if err != nil {
		return
	}

	// send the message
	whatsappClient.Message.Send(manager.SendMessageParams{Message: textMessage, PhoneNumber: "919643500545"})
}
