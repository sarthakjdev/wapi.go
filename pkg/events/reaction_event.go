package events

type ReactionMessageEvent struct {
	BaseMessageEvent
}

func NewReactionMessageEvent(baseMessageEvent BaseMessageEvent, text string) *ReactionMessageEvent {
	return &ReactionMessageEvent{
		BaseMessageEvent: baseMessageEvent,
	}
}
