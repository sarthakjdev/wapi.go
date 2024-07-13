package events

type AccountReviewUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewAccountReviewUpdateEvent() *AccountReviewUpdateEvent {
	return &AccountReviewUpdateEvent{}
}
