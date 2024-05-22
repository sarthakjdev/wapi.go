package manager

import (
	"fmt"

	"github.com/sarthakjdev/wapi.go/pkg/models"
)

type MessageManager struct {
	requester requestClient
}

func NewMessageManager(requester requestClient) *MessageManager {
	return &MessageManager{
		requester: requester,
	}
}

type SendMessageParams struct {
	Message     models.BaseMessage
	PhoneNumber string
}

// ! TODO: return the structured response from here
func (mm *MessageManager) Send(params SendMessageParams) (string, error) {
	body, err := params.Message.ToJson(models.ApiCompatibleJsonConverterConfigs{
		SendToPhoneNumber: params.PhoneNumber,
		// ReplyToMessageId: "wamid.HBgMOTE5NjQzNTAwNTQ1FQIAERgSQzVGOTlFMzExQ0VCQTg0MUFCAA==",
	})
	if err != nil {
		// emit a error event here
		return "", fmt.Errorf("error converting message to json: %v", err)
	}
	mm.requester.requestCloudApi(requestCloudApiParams{
		body: string(body),
		path: "/" + mm.requester.phoneNumberId + "/messages",
	})
	return "ok", nil
}

func (mm *MessageManager) Reply() {
	// Reply to message
}
