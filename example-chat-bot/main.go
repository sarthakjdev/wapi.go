package main

import (
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal/manager"
	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
	wapiComponents "github.com/sarthakjdev/wapi.go/pkg/components"
	"github.com/sarthakjdev/wapi.go/pkg/events"
)

func main() {

	// ! TODO: programmatic read the env variables here

	// creating a client
	whatsappClient, err := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "113269274970227",
		ApiAccessToken:    "EABhCftGVaeIBO27nfwWbjXQWIxFbXFbcRLNbtavOmdGQSZBq7cQ0L6pqACzW2VZBEy53fUohFJQNLdVqeS4hnthQrKS2X9d0Wm1rXih7ej8nkEUjxI0odIq3VLCjlD2RaPK7ZA0BA0lBhkDVmoDbQX7cUIObu6yqUvY2rsDiKZCZA6qNQocjU40e0z6a97kjbt3t7Inkp5R55BSF8uzVEv5zKN0ZBU71Rap7aEoZBplAsjy",
		BusinessAccountId: "103043282674158",
		WebhookPath:       "/webhook",
		WebhookSecret:     "1234567890",
		WebhookServerPort: 8080,
	})

	if err != nil {
		fmt.Println("error creating client", err)
		return
	}

	// create a message
	textMessage, err := wapiComponents.NewTextMessage(wapiComponents.TextMessageConfigs{
		Text: "Hello, from wapi.go",
	})

	if err != nil {
		fmt.Println("error creating text message", err)
		return
	}

	// contactMessage, err := wapiComponents.NewContactMessage([]wapiComponents.Contact{
	// 	*wapiComponents.NewContact(wapiComponents.ContactName{
	// 		FormattedName: "Sarthak Jain",
	// 		FirstName:     "Sarthak",
	// 	})})

	// if err != nil {
	// 	fmt.Println("error creating contact message", err)
	// 	return
	// }

	// locationMessage, err := wapiComponents.NewLocationMessage(28.7041, 77.1025)

	// if err != nil {
	// 	fmt.Println("error creating location message", err)
	// 	return
	// }

	// reactionMessage, err := wapiComponents.NewReactionMessage(wapiComponents.ReactionMessageParams{
	// 	MessageId: "wamid.HBgMOTE5NjQzNTAwNTQ1FQIAERgSQzVGOTlFMzExQ0VCQTg0MUFCAA==",
	// 	Emoji:     "😍",
	// })

	// if err != nil {
	// 	fmt.Println("error creating reaction message", err)
	// 	return
	// }

	// send the message
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: textMessage, PhoneNumber: "919643500545"})
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: contactMessage, PhoneNumber: "919643500545"})
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: locationMessage, PhoneNumber: "919643500545"})
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: reactionMessage, PhoneNumber: "919643500545"})

	// listMessage, err := wapiComponents.NewListMessage(wapiComponents.ListMessageParams{
	// 	ButtonText: "Button 1",
	// 	BodyText:   "Body 1",
	// })

	// if err != nil {
	// 	fmt.Println("error creating list message", err)
	// 	return
	// }

	// listSectionRow, err := wapiComponents.NewListSectionRow("1", "Title 1", "Description 1")

	// if err != nil {
	// 	fmt.Println("error creating list section row", err)
	// 	return
	// }

	// listSection, err := wapiComponents.NewListSection("Section1")

	// if err != nil {
	// 	fmt.Println("error creating list section row", err)
	// 	return
	// }

	// listSection.AddRow(listSectionRow)
	// listMessage.AddSection(listSection)
	// jsonData, err := listMessage.ToJson(wapiComponents.ApiCompatibleJsonConverterConfigs{SendToPhoneNumber: "919643500545"})

	// if err != nil {
	// 	fmt.Println("error converting message to json", err)
	// 	return
	// }

	// fmt.Println(string(jsonData))

	// whatsappClient.Message.Send(manager.SendMessageParams{Message: listMessage, PhoneNumber: "919643500545"})

	// buttonMessage, err := wapiComponents.NewQuickReplyButtonMessage("Body 1")

	// if err != nil {
	// 	fmt.Println("error creating button message", err)
	// 	return
	// }

	// buttonMessage.AddButton("1", "Button 1")
	// buttonMessage.AddButton("2", "Button 2")

	// whatsappClient.Message.Send(manager.SendMessageParams{Message: buttonMessage, PhoneNumber: "919643500545"})

	whatsappClient.On(manager.ReadyEvent, func(event events.BaseEvent) {
		fmt.Println("client is ready")
	})

	whatsappClient.On(manager.TextMessageEvent, func(event events.BaseEvent) {
		fmt.Println("text message event received")

		textMessageEvent := event.(*events.TextMessageEvent)
		fmt.Println(textMessageEvent.Context.From)
		textMessageEvent.Reply(textMessage)

	})

	whatsappClient.On(manager.AudioMessageEvent, func(be events.BaseEvent) {
		fmt.Println("audio message event received")
	})

	whatsappClient.On(manager.VideoMessageEvent, func(be events.BaseEvent) {
		fmt.Println("video message event received")
	})

	whatsappClient.On(manager.DocumentMessageEvent, func(be events.BaseEvent) {
		fmt.Println("document message event received")
	})

	whatsappClient.On(manager.ImageMessageEvent, func(be events.BaseEvent) {
		fmt.Println("image message event received")
	})

	whatsappClient.InitiateClient()
}
