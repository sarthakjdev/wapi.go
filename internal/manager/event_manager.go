package manager

import (
	"fmt"
	"sync"

	"github.com/sarthakjdev/wapi.go/pkg/events"
)

type EventType string

const (
	TextMessageEvent             EventType = "text_message"
	AudioMessageEvent            EventType = "audio_message"
	VideoMessageEvent            EventType = "video_message"
	ImageMessageEvent            EventType = "image_message"
	ContactMessageEvent          EventType = "contact_message"
	DocumentMessageEvent         EventType = "document_message"
	LocationMessageEvent         EventType = "location_message"
	ReactionMessageEvent         EventType = "reaction_message"
	ListInteractionMessageEvent  EventType = "list_interaction_message"
	TemplateMessageEvent         EventType = "template_message"
	QuickReplyMessageEvent       EventType = "quick_reply_message"
	ReplyButtonInteractionEvent  EventType = "reply_button_interaction"
	StickerMessageEvent          EventType = "sticker_message"
	AdInteractionEvent           EventType = "ad_interaction_message"
	CustomerIdentityChangedEvent EventType = "customer_identity_changed"
	CustomerNumberChangedEvent   EventType = "customer_number_changed"
	MessageDeliveredEvent        EventType = "message_delivered"
	MessageFailedEvent           EventType = "message_failed"
	MessageReadEvent             EventType = "message_read"
	MessageSentEvent             EventType = "message_sent"
	MessageUndeliveredEvent      EventType = "message_undelivered"
	OrderReceivedEvent           EventType = "order_received"
	ProductInquiryEvent          EventType = "product_inquiry"
	UnknownEvent                 EventType = "unknown"
	ErrorEvent                   EventType = "error"
	WarnEvent                    EventType = "warn"
	ReadyEvent                   EventType = "ready"
)

type ChannelEvent struct {
	Type EventType
	Data events.BaseEvent
}

type EventManger struct {
	subscribers map[string]chan ChannelEvent
	sync.RWMutex
}

func NewEventManager() *EventManger {
	return &EventManger{
		subscribers: make(map[string]chan ChannelEvent),
	}
}

// subscriber to this event listener will be notified when the event is published
func (em *EventManger) Subscribe(eventName string) (chan ChannelEvent, error) {
	em.Lock()
	defer em.Unlock()
	if ch, ok := em.subscribers[eventName]; ok {
		return ch, nil
	}
	em.subscribers[eventName] = make(chan ChannelEvent, 100)
	return em.subscribers[eventName], nil

}

// subscriber to this event listener will be notified when the event is published
func (em *EventManger) Unsubscribe(id string) {
	em.Lock()
	defer em.Unlock()
	delete(em.subscribers, id)
}

// publish event to this events system and let all the subscriber consume them
func (em *EventManger) Publish(eventType EventType, data events.BaseEvent) error {
	fmt.Println("Publishing event: ", eventType)
	em.Lock()
	defer em.Unlock()

	for _, ch := range em.subscribers {
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

func (em *EventManger) On(name EventType, handler func(events.BaseEvent)) string {
	ch, _ := em.Subscribe(string(name))
	go func() {
		for {
			select {
			case event := <-ch:
				handler(event.Data)
			}
		}
	}()
	return string(name)
}
