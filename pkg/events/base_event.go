package events

type MessageContext struct {
	From string `json:"from"`
}

type BaseEvent interface {
	GetEventType() string
}

type BaseMessageEventInterface interface {
	BaseEvent
	Reply() (string, error)
	React() (string, error)
}

type BaseSystemEventInterface interface {
	BaseEvent
}

type BaseMessageEvent struct {
	MessageId string         `json:"message_id"`
	Context   MessageContext `json:"context"`
}

func NewBaseMessageEvent(messageId, from string) BaseMessageEvent {
	return BaseMessageEvent{
		MessageId: messageId,
		Context: MessageContext{
			From: from,
		},
	}
}

func (bme BaseMessageEvent) GetEventType() string {
	return "message"
}

func (baseMessageEvent *BaseMessageEvent) Reply() (string, error) {
	// we need requester here
	return "", nil

}

func (baseMessageEvent *BaseMessageEvent) React() (string, error) {
	// we need requester here
	return "", nil
}

type BaseSystemEvent struct {
}

func (bme BaseSystemEvent) GetEventType() string {
	return "system"
}
