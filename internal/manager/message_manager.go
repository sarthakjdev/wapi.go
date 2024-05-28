package manager

import (
	"fmt"

	requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"
	"github.com/sarthakjdev/wapi.go/pkg/components"
)

type MessageManager struct {
	requester requestclient.RequestClient
}

func NewMessageManager(requester requestclient.RequestClient) *MessageManager {
	return &MessageManager{
		requester: requester,
	}
}

type SendMessageParams struct {
	Message     components.BaseMessage
	PhoneNumber string
}

// ! TODO: return the structured response from here
func (mm *MessageManager) Send(params SendMessageParams) (string, error) {
	body, err := params.Message.ToJson(components.ApiCompatibleJsonConverterConfigs{
		SendToPhoneNumber: params.PhoneNumber,
		// ReplyToMessageId: "wamid.HBgMOTE5NjQzNTAwNTQ1FQIAERgSQzVGOTlFMzExQ0VCQTg0MUFCAA==",
	})
	if err != nil {
		// ! TODO: emit a error event here
		return "", fmt.Errorf("error converting message to json: %v", err)
	}
	mm.requester.RequestCloudApi(requestclient.RequestCloudApiParams{
		Body: string(body),
		Path: "/" + mm.requester.PhoneNumberId + "/messages",
	})
	return "ok", nil
}

func (mm *MessageManager) Reply() {
	// Reply to message
}
