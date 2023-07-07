package events

import (
	"konn/entity"
)

type UserCreateEvent struct {
	to entity.User
}

func NewUserCreateEvent(to entity.User) Event {
	return UserCreateEvent{
		to: to,
	}
}

func (e UserCreateEvent) toMessage() string {
	return e.String()
}

func (e UserCreateEvent) String() string {
	return BEGIN + br + "EventName: "
}
