package main

import (
	"fmt"

	"github.com/sarthakjdev/wapi.go/internal/manager"
	wapi "github.com/sarthakjdev/wapi.go/pkg/client"
	wapiComponents "github.com/sarthakjdev/wapi.go/pkg/components"
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
	textMessage, err := wapiComponents.NewTextMessage(wapiComponents.TextMessageConfigs{
		Text: "Hello, from wapi.go",
	})

	// if err != nil {
	// 	fmt.Println("error creating text message", err)
	// 	return
	// }

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
	whatsappClient.Message.Send(manager.SendMessageParams{Message: textMessage, PhoneNumber: "919643500545"})
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: contactMessage, PhoneNumber: "919643500545"})
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: locationMessage, PhoneNumber: "919643500545"})
	// whatsappClient.Message.Send(manager.SendMessageParams{Message: reactionMessage, PhoneNumber: "919643500545"})

	listMessage, err := wapiComponents.NewListMessage(wapiComponents.ListMessageParams{
		ButtonText: "Button 1",
		BodyText:   "Body 1",
	})

	if err != nil {
		fmt.Println("error creating list message", err)
		return
	}

	listSectionRow, err := wapiComponents.NewListSectionRow("1", "Title 1", "Description 1")

	if err != nil {
		fmt.Println("error creating list section row", err)
		return
	}

	listSection, err := wapiComponents.NewListSection("Section1")

	if err != nil {
		fmt.Println("error creating list section row", err)
		return
	}

	listSection.AddRow(listSectionRow)
	listMessage.AddSection(listSection)
	jsonData, err := listMessage.ToJson(wapiComponents.ApiCompatibleJsonConverterConfigs{SendToPhoneNumber: "919643500545"})

	if err != nil {
		fmt.Println("error converting message to json", err)
		return
	}

	fmt.Println(string(jsonData))

	whatsappClient.Message.Send(manager.SendMessageParams{Message: listMessage, PhoneNumber: "919643500545"})

}
