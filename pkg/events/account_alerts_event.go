package events

type AccountAlertEvent struct {
	BaseBusinessAccountEvent
}

func NewAccountAlertEvent() *AccountAlertEvent {
	return &AccountAlertEvent{}
}
