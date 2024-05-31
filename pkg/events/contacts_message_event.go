package events

type ContactsMessageEvent struct {
	BaseMessageEvent
}

func NewContactsMessageEvent(baseMessageEvent BaseMessageEvent, text string) *ContactsMessageEvent {
	return &ContactsMessageEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
