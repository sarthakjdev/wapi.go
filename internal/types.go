package internal

type WhatsAppBusinessApiPaginationMeta struct {
	Paging struct {
		Cursors struct {
			Before string `json:"before,omitempty"`
			After  string `json:"after,omitempty"`
		} `json:"cursors,omitempty"`
	} `json:"paging,omitempty"`
}
