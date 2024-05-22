package main

import (
	"github.com/sarthakjdev/wapi.go/internal/manager"
	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
	wapiModels "github.com/sarthakjdev/wapi.go/pkg/models"
)

func main() {
	// creating a client
	whatsappClient := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "113269274970227",
		ApiAccessToken:    "EABhCftGVaeIBOZCmTPNEWyAWSq1Bna8ZCs2Rl7YoXtHfjuXZCAZCMVkjUD8pauKlkGHO6NHP5dXLG9zG5wH7qZAm2GBu8ZChKo37TBa2LzFedZA5sAAZBfgKJ7k7sQZB8t5neJ46DiTH4ZAjxFGvBWRJbC2CP30WHZBhES64pcLCSrWAkoitZBCKAwt5NkOZCoYZCymbP1LnK7e7SZC72R60UJZB4h0xOUvZCF99ohbeO5qeSuxou3Y0ZD",
		BusinessAccountId: "103043282674158",
		WebhookPath:       "/webhook",
		WebhookSecret:     "123456789",
		WebhookServerPort: 8080,
	})
	// create a message
	textMessage := wapiModels.NewTextMessage(wapiModels.TextMessageConfigs{
		Text:         "Hello",
		AllowPreview: true,
	})

	contactMessage := wapiModels.NewContactMessage(wapiModels.ContactMessageConfigs{})

	// send the message
	whatsappClient.Message.Send(manager.SendMessageParams{Message: textMessage, PhoneNumber: "919643500545"})
	whatsappClient.Message.Send(manager.SendMessageParams{Message: contactMessage, PhoneNumber: "919643500545"})
}
