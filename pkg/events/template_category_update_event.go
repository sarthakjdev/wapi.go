package events

type TemplateCategoryUpdateEvent struct {
	BaseBusinessAccountEvent
}

func NewMessageTemplateCategoryUpdateEvent() *TemplateCategoryUpdateEvent {
	return &TemplateCategoryUpdateEvent{}
}
