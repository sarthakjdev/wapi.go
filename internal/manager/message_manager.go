package manager

import (
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
func (mm *MessageManager) Send(params SendMessageParams) {
	// pass the phone number in this toJson method of every message
	body, err := params.Message.ToJson()
	if err != nil {
		// emit a error event here
		return
	}

	mm.requester.requestCloudApi(requestCloudApiParams{
		body: string(body),
		path: "/" + mm.requester.phoneNumberId + "/messages",
	})

	// if err != nil {
	// 	// emit a error event here
	// 	return
	// }

	// fmt.Println("Response from cloud api is", response)

}

func (mm *MessageManager) Reply() {
	// Reply to message
}
