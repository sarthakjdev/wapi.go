package events

type BaseEvent struct {
}

type BaseMessageEvent struct {
	BaseEvent
}

type BaseSystemEvent struct {
	BaseEvent
}

func (*BaseMessageEvent) Reply() {

}

func (*BaseMessageEvent) React() {

}
