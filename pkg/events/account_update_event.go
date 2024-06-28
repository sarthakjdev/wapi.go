package events

type AccountUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewAccountUpdateEvent() *AccountUpdateEvent {
	return &AccountUpdateEvent{}
}
