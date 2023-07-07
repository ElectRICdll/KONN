package events

import (
	"konn/entity"
)

type ColorChangeEvent struct {
	to    *entity.User
	after int
	isMe  bool
}

func NewColorChangeEvent(to *entity.User, after int) Event {
	return ColorChangeEvent{
		to:    to,
		after: after,
		isMe:  true,
	}
}

func (e ColorChangeEvent) toMessage() string {
	return e.String()
}

func (e ColorChangeEvent) String() string {
	return BEGIN + br + "EventName: "
}
