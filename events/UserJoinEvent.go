package events

import (
	"konn/entity"
)

type UserLoginEvent struct {
	which entity.User
}

func NewUserLoginEvent(which entity.User) Event {
	return UserLoginEvent{
		which: which,
	}
}

func (e UserLoginEvent) ToMessage() string {
	return e.String()
}

func (e UserLoginEvent) String() string {
	return BEGIN + br + "EventName: "
}
