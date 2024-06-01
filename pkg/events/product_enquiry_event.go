package events

// ProductInquiryEvent represents an event related to a product inquiry.
type ProductInquiryEvent struct {
	BaseMessageEvent `json:",inline"`
	ProductId        string `json:"productId"`
	CatalogId        string `json:"catalogId"`
	Text             string `json:"text"`
}

// NewProductInquiryEvent creates a new instance of ProductInquiryEvent.
func NewProductInquiryEvent(baseMessageEvent BaseMessageEvent, productId, catalogId, text string) *ProductInquiryEvent {
	return &ProductInquiryEvent{
		BaseMessageEvent: baseMessageEvent,
		ProductId:        productId,
		CatalogId:        catalogId,
		Text:             text,
	}
}
