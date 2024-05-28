package events

type ReadyEvent struct {
	BaseSystemEvent
}

func NewReadyEvent() *ReadyEvent {
	return &ReadyEvent{}
}
