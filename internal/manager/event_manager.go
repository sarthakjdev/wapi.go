package manager

import (
	"fmt"
	"sync"

	"github.com/sarthakjdev/wapi.go/pkg/events"
)

type ChannelEvent struct {
	Type EventType
	Data events.BaseEvent
}

type EventManger struct {
	subscribers map[EventType]chan ChannelEvent
	sync.RWMutex
}

func NewEventManager() *EventManger {
	return &EventManger{
		subscribers: make(map[EventType]chan ChannelEvent),
	}
}

// subscriber to this event listener will be notified when the event is published
func (em *EventManger) Subscribe(eventName EventType) (chan ChannelEvent, error) {
	em.Lock()
	defer em.Unlock()
	if ch, ok := em.subscribers[eventName]; ok {
		return ch, nil
	}
	em.subscribers[eventName] = make(chan ChannelEvent, 100)
	return em.subscribers[eventName], nil
}

// subscriber to this event listener will be notified when the event is published
func (em *EventManger) Unsubscribe(id EventType) {
	em.Lock()
	defer em.Unlock()
	delete(em.subscribers, id)
}

// publish event to this events system and let all the subscriber consume them
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
