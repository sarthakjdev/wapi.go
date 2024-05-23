package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type VideoMessage struct {
	Id      string `json:"id,omitempty"`
	Link    string `json:"link,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type VideoMessageApiPayload struct {
	BaseMessagePayload
	Video VideoMessage `json:"video" validate:"required"`
}

type VideoMessageConfigs = VideoMessage

func NewVideoMessage(params VideoMessageConfigs) (*VideoMessage, error) {
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

	return &VideoMessage{}, nil
}

func (video *VideoMessage) SetCaption(params string) {
	video.Caption = params
}

func (video *VideoMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := VideoMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, VideoMessageType),
		Video:              *video,
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
