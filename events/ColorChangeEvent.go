package events

import (
	"konn/entity/basic"
)

type ColorChangeEvent struct {
	to    *basic.Player
	after int
	isMe  bool
}

func NewColorChangeEvent(to *basic.Player, after int) Event {
	return ColorChangeEvent{
		to:    to,
		after: after,
		isMe:  true,
	}
}

func (e ColorChangeEvent) ToMessage() string {
	return e.String()
}

func (e ColorChangeEvent) String() string {
	return BEGIN + br + "EventName: "
}
