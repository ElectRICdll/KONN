package events

type eventBus struct {
	eventPool []chan Event
	maxNum    int
}
