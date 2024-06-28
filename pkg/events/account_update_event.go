package events

type AccountUpdateEvent struct {
}

func NewAccountUpdateEvent() *AccountUpdateEvent {
	return &AccountUpdateEvent{}
}
