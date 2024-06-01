package events

type CustomerIdentityChangedEvent struct {
	BaseSystemEvent   `json:",inline"`
	Acknowledged      string `json:"acknowledged"`
	CreationTimestamp string `json:"creationTime"`
	Hash              string `json:"hash"`
}
