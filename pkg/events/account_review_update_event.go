package events

type AccountReviewUpdateEvent struct {
}

func NewAccountReviewUpdateEvent() *AccountReviewUpdateEvent {
	return &AccountReviewUpdateEvent{}
}
