package events

import "konn/entity/prop"

type ExposeEvent struct {
	which prop.Substantive
}

func NewExposeEvent(which prop.Substantive) Event {
	return ExposeEvent{
		which: which,
	}
}

func (e ExposeEvent) toMessage() string {
	return e.String()
}

func (e ExposeEvent) String() string {
	return ""
}
