package events

import (
	"konn/entity/basic"
)

type UserCreateEvent struct {
	to basic.Player
}

func NewUserCreateEvent(to basic.Player) Event {
	return UserCreateEvent{
		to: to,
	}
}

func (e UserCreateEvent) ToMessage() string {
	return e.String()
}

func (e UserCreateEvent) String() string {
	return BEGIN + br + "EventName: "
}
