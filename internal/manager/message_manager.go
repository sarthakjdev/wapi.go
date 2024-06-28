package manager

import (
	"fmt"
	"net/http"
	"strings"

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
	})
	if err != nil {
		// TODO: emit an error event here
		return "", fmt.Errorf("error converting message to json: %v", err)
	}

	apiRequest := mm.requester.NewApiRequest(strings.Join([]string{mm.PhoneNumberId, "messages"}, "/"), http.MethodPost)
	apiRequest.SetBody(string(body))
	response, err := apiRequest.Execute()
	return response, err
}
