package events

// AdInteractionSourceType represents the source type of an ad interaction.
type AdInteractionSourceType string

const (
	// AdInteractionSourceTypePost indicates that the ad interaction is from a post.
	AdInteractionSourceTypePost AdInteractionSourceType = "post"
	// AdInteractionSourceTypeAd indicates that the ad interaction is from an ad.
	AdInteractionSourceTypeAd AdInteractionSourceType = "ad"
)

// AdInteractionSourceMediaType represents the media type of an ad interaction source.
type AdInteractionSourceMediaType string

const (
	// AdInteractionSourceMediaTypeImage indicates that the ad interaction source is an image.
	AdInteractionSourceMediaTypeImage AdInteractionSourceMediaType = "image"
	// AdInteractionSourceMediaTypeVideo indicates that the ad interaction source is a video.
	AdInteractionSourceMediaTypeVideo AdInteractionSourceMediaType = "video"
)

// AdSource represents the source of an ad.
type AdSource struct {
	Url          string                       `json:"url"`
	Id           string                       `json:"id"`
	Type         AdInteractionSourceType      `json:"type"`
	Title        string                       `json:"title"`
	Description  string                       `json:"description"`
	MediaUrl     string                       `json:"mediaUrl"`
	MediaType    AdInteractionSourceMediaType `json:"mediaType"`
	ThumbnailUrl string                       `json:"thumbnailUrl"`
	CtwaClid     string                       `json:"ctwaClid"`
}

// AdInteractionEvent represents an ad interaction event.
type AdInteractionEvent struct {
	BaseMessageEvent `json:",inline"`
	AdSource         AdSource `json:"adSource"`
	Text             string   `json:"text"`
}

// NewAdInteractionEvent creates a new instance of AdInteractionEvent.
func NewAdInteractionEvent(baseMessageEvent BaseMessageEvent, adSource AdSource, text string) *AdInteractionEvent {
	return &AdInteractionEvent{
		BaseMessageEvent: baseMessageEvent,
		AdSource:         adSource,
		Text:             text,
	}
}
