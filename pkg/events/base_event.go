package events

import (
	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/components"
)

type MessageContext struct {
	From string `json:"from"`
}

type BaseEvent interface {
	GetEventType() string
}

type BaseMessageEventInterface interface {
	BaseEvent
	Reply() (string, error)
	React() (string, error)
}

type BaseSystemEventInterface interface {
	BaseEvent
}

type BaseMessageEvent struct {
	requester   requestclient.RequestClient
	MessageId   string         `json:"message_id"`
	Context     MessageContext `json:"context"`
	Timestamp   string         `json:"timestamp"`
	IsForwarded bool           `json:"is_forwarded"`
}

func NewBaseMessageEvent(messageId string, timestamp string, from string, isForwarded bool, requester requestclient.RequestClient) BaseMessageEvent {
	return BaseMessageEvent{
		MessageId: messageId,
		Context: MessageContext{
			From: from,
		},
		requester:   requester,
		Timestamp:   timestamp,
		IsForwarded: isForwarded,
	}
}

func (bme BaseMessageEvent) GetEventType() string {
	return "message"
}

// Reply to the message
func (baseMessageEvent *BaseMessageEvent) Reply(Message components.BaseMessage) (string, error) {

	body, err := Message.ToJson(components.ApiCompatibleJsonConverterConfigs{
		SendToPhoneNumber: baseMessageEvent.Context.From,
		ReplyToMessageId:  baseMessageEvent.MessageId,
	})

	if err != nil {
		return "", err
	}

	baseMessageEvent.requester.RequestCloudApi(requestclient.RequestCloudApiParams{
		Body: string(body),
		Path: "/" + baseMessageEvent.requester.PhoneNumberId + "/messages",
	})

	return "", nil

}

// React to the message
func (baseMessageEvent *BaseMessageEvent) React(emoji string) (string, error) {
	reactionMessage, err := components.NewReactionMessage(components.ReactionMessageParams{
		Emoji:     emoji,
		MessageId: baseMessageEvent.MessageId,
	})
	if err != nil {
		return "", err
	}
	baseMessageEvent.Reply(reactionMessage)
	return "", nil
}

type BaseSystemEvent struct {
}

func (bme BaseSystemEvent) GetEventType() string {
	return "system"
}

