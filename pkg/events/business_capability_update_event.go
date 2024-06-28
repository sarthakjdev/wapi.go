package events

type BusinessCapabilityUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewBusinessCapabilityUpdateEvent() *BusinessCapabilityUpdateEvent {
	return &BusinessCapabilityUpdateEvent{}
}
