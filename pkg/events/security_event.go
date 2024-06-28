package events

type SecurityEvent struct {
	BaseBusinessAccountEvent
}

func NewSecurity() *SecurityEvent {
	return &SecurityEvent{}
}
