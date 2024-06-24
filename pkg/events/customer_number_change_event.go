package events

type CustomerNumberChangedEvent struct {
	BaseSystemEvent   `json:",inline"`
	ChangeDescription string `json:"changeDescription"`
	NewWaId           string `json:"newWaId"`
	OldWaId           string `json:"oldWaId"`
}
