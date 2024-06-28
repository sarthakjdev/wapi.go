package events

type PhoneNumberNameUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewPhoneNumberNameUpdateEvent() *PhoneNumberNameUpdateEvent {
	return &PhoneNumberNameUpdateEvent{}
}
