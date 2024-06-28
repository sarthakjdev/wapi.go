package events

type PhoneNumberQualityUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewPhoneNumberQualityUpdateEvent() *PhoneNumberQualityUpdateEvent {
	return &PhoneNumberQualityUpdateEvent{}
}
