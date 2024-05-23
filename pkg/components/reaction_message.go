package components

type ReactionMessage struct {
}

func NewReactionMessage() (*ReactionMessage, error) {
	return &ReactionMessage{}, nil
}

func (*ReactionMessage) ToJson() {

}



