package events

import (
	"fmt"
	"konn/ingame/Players"
	"reflect"
)

type NewUserEvent struct {
	*Event
	NewUser *Players.User
	WHETHER_ME  bool
}

func (e *NewUserEvent) Registry(newUser *Players.User) {
	e.NewUser = newUser
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<new_user:%s>", e.NewUser.Name)
	e.Event = NewEvent(newUser.Name, reflect.TypeOf(e).String(), msg)
}

func (e *NewUserEvent) String() string {
	return BEGIN + br + "EventName: " + reflect.TypeOf(e).String() + br + e.Event.Message
}
