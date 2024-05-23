package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type LocationMessage struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Address   string  `json:"address,omitempty"`
	Name      string  `json:"name,omitempty"`
}

type LocationMessageApiPayload struct {
	BaseMessagePayload
	Location LocationMessage `json:"location" validate:"required"`
}

func NewLocationMessage(latitude float64, longitude float64) (*LocationMessage, error) {

	return &LocationMessage{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil

}

func (location *LocationMessage) SetAddress(params string) {
	location.Address = params
}

func (location *LocationMessage) SetName(params string) {
	location.Name = params
}

func (location *LocationMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := LocationMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeLocation),
		Location:           *location,
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
