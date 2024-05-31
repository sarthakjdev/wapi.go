package manager

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

type NotificationReasonEnum string

const (
	NotificationReasonMessage NotificationReasonEnum = "message"
)

type NotificationPayloadErrorSchemaType struct {
	Code      int    `json:"code"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	ErrorData struct {
		Details string `json:"details"`
	} `json:"error_data,omitempty"`
}

type NotificationPayloadMessageContextSchemaType struct {
	Forwarded           bool   `json:"forwarded,omitempty"`
	FrequentlyForwarded bool   `json:"frequently_forwarded,omitempty"`
	From                string `json:"from,omitempty"`
	Id                  string `json:"id"`
	ReferredProduct     struct {
		CatalogId         string `json:"catalog_id"`
		ProductRetailerId string `json:"product_retailer_id"`
	} `json:"referred_product,omitempty"`
}

type NotificationPayloadTextMessageSchemaType struct {
	Text struct {
		Body string `json:"body"`
	} `json:"text,omitempty"`
	Referral struct {
		SourceUrl    string                           `json:"source_url"`
		SourceType   AdInteractionSourceTypeEnum      `json:"source_type"`
		SourceId     string                           `json:"source_id"`
		Headline     string                           `json:"headline"`
		Body         string                           `json:"body"`
		ImageUrl     string                           `json:"image_url,omitempty"`
		VideoUrl     string                           `json:"video_url,omitempty"`
		ThumbnailUrl string                           `json:"thumbnail_url"`
		CtwaCLId     string                           `json:"ctwa_clid"`
		MediaType    AdInteractionSourceMediaTypeEnum `json:"media_type"`
	} `json:"referral,omitempty"`
}

type NotificationPayloadAudioMessageSchemaType struct {
	Audio struct {
		Id       string `json:"id,omitempty"`
		MIMEType string `json:"mime_type,omitempty"`
		SHA256   string `json:"sha256,omitempty"`
	} `json:"audio,omitempty"`
}

type NotificationPayloadImageMessageSchemaType struct {
	Image struct {
		Id       string `json:"id"`
		MIMEType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		Caption  string `json:"caption,omitempty"`
	} `json:"image,omitempty"`
}

type NotificationPayloadButtonMessageSchemaType struct {
	Button struct {
		Payload string `json:"payload"`
		Text    string `json:"text"`
	} `json:"button,omitempty"`
}

type NotificationPayloadDocumentMessageSchemaType struct {
	Document struct {
		Id       string `json:"id"`
		MIMEType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		Caption  string `json:"caption,omitempty"`
		Filename string `json:"filename,omitempty"`
	} `json:"document,omitempty"`
}

type NotificationPayloadOrderMessageSchemaType struct {
	// OrderText string `json:"text"`
	Order struct {
		CatalogId    string `json:"catalog_id"`
		ProductItems []struct {
			ProductRetailerId string `json:"product_retailer_id"`
			Quantity          string `json:"quantity"`
			ItemPrice         string `json:"item_price"`
			Currency          string `json:"currency"`
		} `json:"product_items"`
	} `json:"order,omitempty"`
}

type NotificationPayloadStickerMessageSchemaType struct {
	Sticker struct {
		Id       string `json:"id"`
		MIMEType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		Animated bool   `json:"animated"`
	} `json:"sticker,omitempty"`
}

type NotificationPayloadSystemMessageSchemaType struct {
	System struct {
		Identity string                     `json:"identity"`
		Body     string                     `json:"body"`
		Customer string                     `json:"customer"`
		Type     SystemNotificationTypeEnum `json:"type"`
		WaId     string                     `json:"wa_id"`
	} `json:"system,omitempty"`
	Identity struct {
		Acknowledged     string `json:"acknowledged"`
		CreatedTimestamp string `json:"created_timestamp"`
		Hash             string `json:"hash"`
	} `json:"identity,omitempty"`
}

type NotificationPayloadVideoMessageSchemaType struct {
	Video struct {
		Id       string `json:"id"`
		MIMEType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
		Caption  string `json:"caption,omitempty"`
		Filename string `json:"filename,omitempty"`
	} `json:"video,omitempty"`
}

type NotificationPayloadReactionMessageSchemaType struct {
	Reaction struct {
		MessageId string `json:"message_id"`
		Emoji     string `json:"emoji"`
	} `json:"reaction,omitempty"`
}

type NotificationPayloadInteractionMessageSchemaType struct {
	Interactive struct {
		Type InteractiveNotificationTypeEnum `json:"type"`
	} `json:"interactive,omitempty"`
}

type NotificationPayloadLocationMessageSchemaType struct {
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Name      string  `json:"name,omitempty"`
		Address   string  `json:"address,omitempty"`
	} `json:"location,omitempty"`
}

type NotificationPayloadContactMessageSchemaType struct {
	Contacts []Contact `json:"contacts"`
}

type NotificationMessageTypeEnum string

const (
	NotificationMessageTypeText        NotificationMessageTypeEnum = "text"
	NotificationMessageTypeAudio       NotificationMessageTypeEnum = "audio"
	NotificationMessageTypeImage       NotificationMessageTypeEnum = "image"
	NotificationMessageTypeButton      NotificationMessageTypeEnum = "button"
	NotificationMessageTypeDocument    NotificationMessageTypeEnum = "document"
	NotificationMessageTypeOrder       NotificationMessageTypeEnum = "order"
	NotificationMessageTypeSticker     NotificationMessageTypeEnum = "sticker"
	NotificationMessageTypeSystem      NotificationMessageTypeEnum = "system"
	NotificationMessageTypeVideo       NotificationMessageTypeEnum = "video"
	NotificationMessageTypeReaction    NotificationMessageTypeEnum = "reaction"
	NotificationMessageTypeInteractive NotificationMessageTypeEnum = "interactive"
	NotificationMessageTypeUnknown     NotificationMessageTypeEnum = "unknown"
	NotificationMessageTypeLocation    NotificationMessageTypeEnum = "location"
	NotificationMessageTypeContacts    NotificationMessageTypeEnum = "contacts"
)

type InteractiveNotificationTypeEnum string

const (
	NotificationTypeButtonReply InteractiveNotificationTypeEnum = "button_reply"
	NotificationTypeListReply   InteractiveNotificationTypeEnum = "list_reply"
)

type AdInteractionSourceTypeEnum string

const (
	AdInteractionSourceTypeUnknown AdInteractionSourceTypeEnum = "unknown"
	// Add other ad interaction source types
)

type AdInteractionSourceMediaTypeEnum string

const (
	AdInteractionSourceMediaTypeImage AdInteractionSourceMediaTypeEnum = "image"
	AdInteractionSourceMediaTypeVideo AdInteractionSourceMediaTypeEnum = "video"
	// Add other ad interaction source media types
)

type SystemNotificationTypeEnum string

const (
	SystemNotificationTypeUnknown SystemNotificationTypeEnum = "unknown"
	// Add other system notification types
)

type Contact struct {
	WaId    string  `json:"wa_id"`
	Profile Profile `json:"profile"`
}

type Profile struct {
	Name string `json:"name"`
}

type WhatsappApiNotificationPayloadSchemaType struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	Id      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string    `json:"messaging_product"`
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts,omitempty"`
	Statuses         []Status  `json:"statuses,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
	Errors           []Error   `json:"errors,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberId      string `json:"phone_number_id"`
}

type Status struct {
	Conversation Conversation `json:"conversation,omitempty"`
	Errors       []Error      `json:"errors,omitempty"`
	Status       string       `json:"status"`
	Timestamp    string       `json:"timestamp"`
	RecipientId  string       `json:"recipient_id"`
	Pricing      Pricing      `json:"pricing,omitempty"`
}

type Conversation struct {
	Id     string `json:"id"`
	Origin Origin `json:"origin,omitempty"`
}

type Origin struct {
	Type                MessageStatusCategoryEnum `json:"type"`
	ExpirationTimestamp string                    `json:"expiration_timestamp,omitempty"`
}

type Pricing struct {
	PricingModel string                    `json:"pricing_model"`
	Category     MessageStatusCategoryEnum `json:"category"`
}

type Message struct {
	Id                                           string                                      `json:"id"`
	From                                         string                                      `json:"from"`
	Timestamp                                    string                                      `json:"timestamp"`
	Type                                         NotificationMessageTypeEnum                 `json:"type"`
	Context                                      NotificationPayloadMessageContextSchemaType `json:"context"`
	Errors                                       []Error                                     `json:",inline"`
	NotificationPayloadTextMessageSchemaType     `json:",inline"`
	NotificationPayloadAudioMessageSchemaType    `json:",inline"`
	NotificationPayloadImageMessageSchemaType    `json:",inline"`
	NotificationPayloadButtonMessageSchemaType   `json:",inline"`
	NotificationPayloadDocumentMessageSchemaType `json:",inline"`
	NotificationPayloadOrderMessageSchemaType    `json:",inline"`
	NotificationPayloadStickerMessageSchemaType  `json:",inline"`
	NotificationPayloadSystemMessageSchemaType   `json:",inline"`
	NotificationPayloadVideoMessageSchemaType    `json:",inline"`
	NotificationPayloadReactionMessageSchemaType `json:",inline"`
	NotificationPayloadLocationMessageSchemaType `json:",inline"`
	NotificationPayloadContactMessageSchemaType  `json:",inline"`
}

type Error struct {
	// Add fields for error details
}

type MessageStatusCategoryEnum string

const (
	MessageStatusCategorySent MessageStatusCategoryEnum = "sent"
)

// Add other message status categories

type MessageStatusEnum string

const (
	MessageStatusDelivered MessageStatusEnum = "delivered"
	// Add other message statuses
)
