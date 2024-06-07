package main

import (
	"fmt"
	"strings"

	wapiComponents "github.com/sarthakjdev/wapi.go/pkg/components"
	"github.com/sarthakjdev/wapi.go/pkg/events"
	wapi "github.com/sarthakjdev/wapi.go/pkg/messaging"
)

func main() {
	// creating a client
	whatsappClient, err := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "",
		ApiAccessToken:    "",
		BusinessAccountId: "",
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

	audioMessage, err := wapiComponents.NewAudioMessage(wapiComponents.AudioMessageConfigs{
		Link: "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
	})

	if err != nil {
		fmt.Println("error creating audio message message", err)
		return
	}

	videoMessage, err := wapiComponents.NewVideoMessage(wapiComponents.VideoMessageConfigs{
		Link: "https://www.sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4",
	})

	if err != nil {
		fmt.Println("error creating video message", err)
		return
	}

	imageMessage, err := wapiComponents.NewImageMessage(wapiComponents.ImageMessageConfigs{
		Link: "https://www.w3schools.com/w3css/img_lights.jpg",
	})

	if err != nil {
		fmt.Println("error creating image message", err)
		return
	}

	contactMessage, err := wapiComponents.NewContactMessage([]wapiComponents.Contact{
		*wapiComponents.NewContact(wapiComponents.ContactName{
			FormattedName: "Sarthak Jain",
			FirstName:     "Sarthak",
		})})

	if err != nil {
		fmt.Println("error creating contact message", err)
		return
	}

	locationMessage, err := wapiComponents.NewLocationMessage(28.7041, 77.1025)

	if err != nil {
		fmt.Println("error creating location message", err)
		return
	}

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

	whatsappClient.Message.Send(listMessage, "919643500545")

	buttonMessage, err := wapiComponents.NewQuickReplyButtonMessage("Body 1")

	if err != nil {
		fmt.Println("error creating button message", err)
		return
	}

	buttonMessage.AddButton("1", "Button 1")
	buttonMessage.AddButton("2", "Button 2")

	whatsappClient.On(events.ReadyEventType, func(event events.BaseEvent) {
		fmt.Println("client is ready")
	})

	whatsappClient.On(events.TextMessageEventType, func(event events.BaseEvent) {
		fmt.Println("text message event received")

		textMessageEvent := event.(*events.TextMessageEvent)

		switch strings.ToLower(textMessageEvent.Text) {
		case "text":
			textMessageEvent.Reply(textMessage)
		case "audio":
			textMessageEvent.Reply(audioMessage)
		case "image":
			textMessageEvent.Reply(imageMessage)
		case "video":
			textMessageEvent.Reply(videoMessage)
		case "contact":
			textMessageEvent.Reply(contactMessage)
		case "location":
			textMessageEvent.Reply(locationMessage)
		case "reaction":
			textMessageEvent.React("😍")
		case "list":
			textMessageEvent.Reply(listMessage)
		case "button":
			textMessageEvent.Reply(buttonMessage)
		default:
			textMessageEvent.Reply(textMessage)
		}
	})

	whatsappClient.On(events.AudioMessageEventType, func(be events.BaseEvent) {
		fmt.Println("audio message event received")
	})

	whatsappClient.On(events.VideoMessageEventType, func(be events.BaseEvent) {
		fmt.Println("video message event received")
	})

	whatsappClient.On(events.DocumentMessageEventType, func(be events.BaseEvent) {
		fmt.Println("document message event received")
	})

	whatsappClient.On(events.ImageMessageEventType, func(be events.BaseEvent) {
		fmt.Println("image message event received")
	})

	whatsappClient.InitiateClient()
}
