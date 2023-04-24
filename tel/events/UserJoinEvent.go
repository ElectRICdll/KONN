package events

import (
	"fmt"
	"konn/ingame/Players"
	"reflect"
)

type UserJoinEvent struct {
	Event
	currentUser *Players.User
	WHETHER_ME  bool
}

func (e *UserJoinEvent) Registry(curUser *Players.User) {
	e.currentUser = curUser
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<user:%s>", e.currentUser.Name)
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}

func (e *UserJoinEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message
}
