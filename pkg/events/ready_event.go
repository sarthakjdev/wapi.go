package events

// ReadyEvent represents an event that is triggered when the system is ready.
type ReadyEvent struct {
	BaseSystemEvent `json:",inline"`
}

// NewReadyEvent creates a new instance of ReadyEvent.
func NewReadyEvent() *ReadyEvent {
	return &ReadyEvent{}
}
