package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type ImageMessage struct {
	Id      string `json:"id,omitempty"`
	Link    string `json:"link,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type ImageMessageApiPayload struct {
	BaseMessagePayload
	Image ImageMessage `json:"image" validate:"required"`
}

type ImageMessageConfigs = ImageMessage

func (image *ImageMessage) SetCaption(params string) {
	image.Caption = params
}

func NewImageMessage(params ImageMessageConfigs) (*ImageMessage, error) {
	if err := utils.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	idSet := params.Id != ""
	linkSet := params.Link != ""

	if idSet && linkSet {
		return nil, fmt.Errorf("only one of ID or Link can be provided")
	}
	if !idSet && !linkSet {
		return nil, fmt.Errorf("either ID or Link must be provided")
	}

	return &ImageMessage{
		Id:      params.Id,
		Link:    params.Link,
		Caption: params.Caption,
	}, nil
}

func (image *ImageMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := ImageMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeImage),
		Image:              *image,
	}

	if configs.ReplyToMessageId != "" {
		jsonData.Context = &Context{
			MessageId: configs.ReplyToMessageId,
		}
	}

	jsonToReturn, err := json.Marshal(jsonData)

	if err != nil {
		return nil, fmt.Errorf("error marshalling json: %v", err)
	}

	return jsonToReturn, nil

}
