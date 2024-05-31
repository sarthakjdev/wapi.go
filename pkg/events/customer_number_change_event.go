package events

type CustomerNumberChangedEvent struct {
}

func NewCustomerNumberChangedEvent(baseMessageEvent BaseMessageEvent, text string) *CustomerNumberChangedEvent {
	return &CustomerNumberChangedEvent{}
}
