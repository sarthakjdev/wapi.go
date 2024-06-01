package manager

import (
	"fmt"

	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/components"
)

// MessageManager is responsible for managing messages.
type MessageManager struct {
	requester requestclient.RequestClient
}

// NewMessageManager creates a new instance of MessageManager.
func NewMessageManager(requester requestclient.RequestClient) *MessageManager {
	return &MessageManager{
		requester: requester,
	}
}

// SendMessageParams represents the parameters for sending a message.
type SendMessageParams struct {
	Message     components.BaseMessage
	PhoneNumber string
}

// Send sends a message with the given parameters and returns the response.
// TODO: return the structured response from here
func (mm *MessageManager) Send(params SendMessageParams) (string, error) {
	body, err := params.Message.ToJson(components.ApiCompatibleJsonConverterConfigs{
		SendToPhoneNumber: params.PhoneNumber,
		// ReplyToMessageId: "wamid.HBgMOTE5NjQzNTAwNTQ1FQIAERgSQzVGOTlFMzExQ0VCQTg0MUFCAA==",
	})
	if err != nil {
		// TODO: emit an error event here
		return "", fmt.Errorf("error converting message to json: %v", err)
	}
	mm.requester.RequestCloudApi(requestclient.RequestCloudApiParams{
		Body: string(body),
		Path: "/" + mm.requester.PhoneNumberId + "/messages",
	})
	return "ok", nil
}

// Reply replies to a message.
func (mm *MessageManager) Reply() {
	// Reply to message
}
