package events

type CustomerIdentityChangedEvent struct {
}

func NewCustomerIdentityChangedEvent(baseMessageEvent BaseMessageEvent, text string) *CustomerIdentityChangedEvent {
	return &CustomerIdentityChangedEvent{}
}
