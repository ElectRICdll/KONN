package events

import (
	"konn/entity"
)

type UserCreateEvent struct {
	to entity.Player
}

func NewUserCreateEvent(to entity.Player) Event {
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
