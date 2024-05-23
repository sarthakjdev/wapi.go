package components

import (
	"encoding/json"
	"fmt"

	"github.com/sarthakjdev/wapi.go/utils"
)

type listSection struct {
	Title string           `json:"title" validate:"required"`
	Rows  []listSectionRow `json:"rows" validate:"required"`
}

func NewListSection(title string) (*listSection, error) {
	return &listSection{
		Title: title,
	}, nil
}

type listSectionRow struct {
	Id          string `json:"id" validate:"required"`
	Description string `json:"description" validate:"required"`
	Title       string `json:"title" validate:"required"`
}

func (section *listSection) AddRow(row *listSectionRow) {
	section.Rows = append(section.Rows, *row)
}

func (section *listSection) SetTitle(title string) {
	section.Title = title
}

func NewListSectionRow(id, title, description string) (*listSectionRow, error) {
	return &listSectionRow{
		Id:          id,
		Description: description,
		Title:       title,
	}, nil
}

func (row *listSectionRow) SetTitle(title string) {
	row.Title = title
}

func (row *listSectionRow) SetDescription(description string) {
	row.Description = description
}

func (row *listSectionRow) SetId(id string) {
	row.Id = id
}

type listMessageAction struct {
	ButtonText string        `json:"button" validate:"required"`
	Sections   []listSection `json:"sections" validate:"required"`
}

type ListMessageBody struct {
	Text string `json:"text" validate:"required"`
}

type listMessage struct {
	Type   InteractiveMessageType `json:"type" validate:"required"`
	Action listMessageAction      `json:"action" validate:"required"`
	Body   ListMessageBody        `json:"body,omitempty"`
}

type ListMessageParams struct {
	ButtonText string `json:"-" validate:"required"`
	BodyText   string `json:"-" validate:"required`
}

func NewListMessage(params ListMessageParams) (*listMessage, error) {
	if err := utils.GetValidator().Struct(params); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	return &listMessage{
		Type: InteractiveMessageTypeList,
		Body: ListMessageBody{
			Text: params.BodyText,
		},
		Action: listMessageAction{
			ButtonText: params.ButtonText,
		},
	}, nil
}

type ListMessageApiPayload struct {
	BaseMessagePayload
	Interactive listMessage `json:"interactive" validate:"required"`
}

func (m *listMessage) AddSection(section *listSection) {
	m.Action.Sections = append(m.Action.Sections, *section)
}

func (m *listMessage) SetBodyText(section *listSection) {
	m.Body.Text = section.Title
}

func (m *listMessage) ToJson(configs ApiCompatibleJsonConverterConfigs) ([]byte, error) {
	if err := utils.GetValidator().Struct(configs); err != nil {
		return nil, fmt.Errorf("error validating configs: %v", err)
	}

	jsonData := ListMessageApiPayload{
		BaseMessagePayload: NewBaseMessagePayload(configs.SendToPhoneNumber, MessageTypeInteractive),
		Interactive:        *m,
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
