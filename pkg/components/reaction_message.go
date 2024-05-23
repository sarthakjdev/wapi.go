package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type ReactionMessage struct {
	MessageId string `json:"message_id" validate:"required"`
	Emoji     string `json:"emoji" validate:"required"`
}

type ReactionMessageParams = ReactionMessage

type ReactionMessageApiPayload struct {
	BaseMessagePayload
	Reaction ReactionMessage `json:"reaction" validate:"required"`
}

func NewReactionMessage(params ReactionMessageParams) (*ReactionMessage, error) {

	if err := utils.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	return &ReactionMessage{
		MessageId: params.MessageId,
		Emoji:     params.Emoji,
	}, nil
}

func (reaction *ReactionMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := ReactionMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeReaction),
		Reaction:           *reaction,
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
