package main

import (
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal/manager"
	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
	wapiModels "github.com/sarthakjdev/wapi.go/pkg/models"
)

func main() {

	// ! TODO: programmatic read the env variables here

	// creating a client
	whatsappClient, err := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "113269274970227",
		ApiAccessToken:    "EABhCftGVaeIBOZCgZCHPf8eF7ZBayGCyVLvpGVbZC8oqjgZCzmhqVXn7TMiQ3JTQ77WxOE4K7DVIgFC8ZA7qSG2ANHQ3BbG09iXezHDHnu2iiC0K5VVcITzHZCMoy5aKkLhILxLNsOQ5s9nQg3dRj1VewJ1PuMJY2n9tcIP29qn0Ht30fpUirvG6tgE9CVdRlMHuHU54U4hFjqcNfbO4Q8jW1QvhKCZBv95do3YFd71v1ucZD",
		BusinessAccountId: "103043282674158",
		WebhookPath:       "/webhook",
		WebhookSecret:     "123456789",
		WebhookServerPort: 8080,
	})

	if err != nil {
		fmt.Println("error creating client", err)
		return
	}

	// create a message
	textMessage, err := wapiModels.NewTextMessage(wapiModels.TextMessageConfigs{
		Text: "Hello, from wapi.go",
	})

	if err != nil {
		fmt.Println("error creating text message", err)
		return
	}

	contact := wapiModels.NewContact(wapiModels.ContactName{
		FormattedName: "Sarthak Jain",
		FirstName:     "Sarthak",
	})

	contactMessage, err := wapiModels.NewContactMessage([]wapiModels.Contact{*contact})

	if err != nil {
		fmt.Println("error creating contact message", err)
		return
	}

	if err != nil {
		fmt.Println("error creating text message", err)
		return
	}

	// send the message
	whatsappClient.Message.Send(manager.SendMessageParams{Message: textMessage, PhoneNumber: "919643500545"})
	whatsappClient.Message.Send(manager.SendMessageParams{Message: contactMessage, PhoneNumber: "919643500545"})
}
