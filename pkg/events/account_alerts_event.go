package events

type AccountAlertEvent struct {
}

func NewAccountAlertEvent() *AccountAlertEvent {
	return &AccountAlertEvent{}
}
