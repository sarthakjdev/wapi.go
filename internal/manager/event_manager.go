package manager

import (
	"fmt"
	"sync"

	"github.com/sarthakjdev/wapi.go/pkg/events"
)

// ChannelEvent represents an event that can be published and subscribed to.
type ChannelEvent struct {
	Type EventType        // Type is the type of the event.
	Data events.BaseEvent // Data is the data associated with the event.
}

// EventManger is responsible for managing events and their subscribers.
type EventManger struct {
	subscribers  map[EventType]chan ChannelEvent // subscribers is a map of event types to channels of ChannelEvent.
	sync.RWMutex                                 // RWMutex is used to synchronize access to the subscribers map.
}

// NewEventManager creates a new instance of EventManger.
func NewEventManager() *EventManger {
	return &EventManger{
		subscribers: make(map[EventType]chan ChannelEvent),
	}
}

// Subscribe adds a new subscriber to the specified event type.
// The subscriber will be notified when the event is published.
func (em *EventManger) Subscribe(eventName EventType) (chan ChannelEvent, error) {
	em.Lock()
	defer em.Unlock()
	if ch, ok := em.subscribers[eventName]; ok {
		return ch, nil
	}
	em.subscribers[eventName] = make(chan ChannelEvent, 100)
	return em.subscribers[eventName], nil
}

// Unsubscribe removes a subscriber from the specified event type.
func (em *EventManger) Unsubscribe(id EventType) {
	em.Lock()
	defer em.Unlock()
	delete(em.subscribers, id)
}

// Publish publishes an event to the event system and notifies all the subscribers.
func (em *EventManger) Publish(eventType EventType, data events.BaseEvent) error {
	fmt.Println("Publishing event: ", eventType)
	em.Lock()
	defer em.Unlock()

	if ch, ok := em.subscribers[eventType]; ok {
		select {
		case ch <- ChannelEvent{
			Type: eventType,
			Data: data,
		}:
		default:
			return fmt.Errorf("event queue full for type: %s", eventType)
		}
	}
	return nil
}

// On registers a handler function for the specified event type.
// The handler function will be called whenever the event is published.
// It returns the event type that the handler is registered for.
func (em *EventManger) On(eventName EventType, handler func(events.BaseEvent)) EventType {
	ch, _ := em.Subscribe(eventName)
	go func() {
		for {
			select {
			case event := <-ch:
				handler(event.Data)
			}
		}
	}()
	return eventName
}
