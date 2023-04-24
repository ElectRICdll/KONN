package events

import (
	"fmt"
	"konn/ingame/Players"
	"reflect"
)

type UserLeftEvent struct {
	*Event
	currentUser *Players.User
	WHETHER_ME  bool
}

func (e *UserLeftEvent) Registry(curUser *Players.User) {
	e.currentUser = curUser  
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<user:%s>", e.currentUser.Name)
	e.Event = NewEvent(curUser.Name, reflect.TypeOf(e).String(), msg)
}

func (e *UserLeftEvent) String() string {
	return BEGIN + br + "EventName: " + reflect.TypeOf(e).String() + br + e.Message
}
