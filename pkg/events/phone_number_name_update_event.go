package events

type PhoneNumberNameUpdateEvent struct{}

func NewPhoneNumberNameUpdateEvent() *PhoneNumberNameUpdateEvent {
	return &PhoneNumberNameUpdateEvent{}
}
