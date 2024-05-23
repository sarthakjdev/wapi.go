package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type AudioMessage struct {
	Id   string `json:"id,omitempty"`
	Link string `json:"link,omitempty"`
}

type AudioMessageApiPayload struct {
	BaseMessagePayload
	Audio AudioMessage `json:"audio" validate:"required"`
}

type AudioMessageConfigs = AudioMessage

func NewAudioMessage(params AudioMessageConfigs) (*AudioMessage, error) {
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

	return &AudioMessage{
		Id:   params.Id,
		Link: params.Link,
	}, nil
}

func (audio *AudioMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := AudioMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeAudio),
		Audio:              *audio,
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
