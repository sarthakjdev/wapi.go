package manager

import (
	"fmt"
	"net/http"

	"github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/components"
)

// MessageManager is responsible for managing messages.
type MessageManager struct {
	requester     request_client.RequestClient
	PhoneNumberId string
}

// NewMessageManager creates a new instance of MessageManager.
func NewMessageManager(requester request_client.RequestClient, phoneNumberId string) *MessageManager {
	return &MessageManager{
		requester:     requester,
		PhoneNumberId: phoneNumberId,
	}
}

// Send sends a message with the given parameters and returns the response.
// TODO: return the structured response from here
func (mm *MessageManager) Send(message components.BaseMessage, phoneNumber string) (string, error) {
	body, err := message.ToJson(components.ApiCompatibleJsonConverterConfigs{
		SendToPhoneNumber: phoneNumber,
		// ReplyToMessageId: "wamid.HBgMOTE5NjQzNTAwNTQ1FQIAERgSQzVGOTlFMzExQ0VCQTg0MUFCAA==",
	})
	if err != nil {
		// TODO: emit an error event here
		return "", fmt.Errorf("error converting message to json: %v", err)
	}
	mm.requester.Request(request_client.RequestCloudApiParams{
		Body:   string(body),
		Path:   "/" + mm.PhoneNumberId + "/messages",
		Method: http.MethodPost,
	})
	return "ok", nil
}
