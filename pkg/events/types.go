package events

type EventType string

const (
	TextMessageEventType             EventType = "text_message"
	AudioMessageEventType            EventType = "audio_message"
	VideoMessageEventType            EventType = "video_message"
	ImageMessageEventType            EventType = "image_message"
	ContactMessageEventType          EventType = "contact_message"
	DocumentMessageEventType         EventType = "document_message"
	LocationMessageEventType         EventType = "location_message"
	ReactionMessageEventType         EventType = "reaction_message"
	ListInteractionMessageEventType  EventType = "list_interaction_message"
	TemplateMessageEventType         EventType = "template_message"
	QuickReplyMessageEventType       EventType = "quick_reply_message"
	ReplyButtonInteractionEventType  EventType = "reply_button_interaction"
	StickerMessageEventType          EventType = "sticker_message"
	AdInteractionEventType           EventType = "ad_interaction_message"
	CustomerIdentityChangedEventType EventType = "customer_identity_changed"
	CustomerNumberChangedEventType   EventType = "customer_number_changed"
	MessageDeliveredEventType        EventType = "message_delivered"
	MessageFailedEventType           EventType = "message_failed"
	MessageReadEventType             EventType = "message_read"
	MessageSentEventType             EventType = "message_sent"
	MessageUndeliveredEventType      EventType = "message_undelivered"
	OrderReceivedEventType           EventType = "order_received"
	ProductInquiryEventType          EventType = "product_inquiry"
	UnknownEventType                 EventType = "unknown"
	ErrorEventType                   EventType = "error"
	WarnEventType                    EventType = "warn"
	ReadyEventType                   EventType = "ready"
)
