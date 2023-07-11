package events

import (
	"konn/entity"
)

type UserLogoutEvent struct {
	which *entity.User
}

func (e UserLogoutEvent) Registry(which *entity.User) Event {
	return UserLogoutEvent{
		which: which,
	}
}

func (e UserLogoutEvent) ToMessage() string {
	return e.String()
}

func (e UserLogoutEvent) String() string {
	return BEGIN + br + "EventName: "
}
