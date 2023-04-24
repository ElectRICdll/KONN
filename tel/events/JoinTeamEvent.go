package events

import (
	"fmt"
	"konn/ingame/Players"
	"reflect"
)

type JoinTeamEvent struct {
	Event
	currentUser *Players.User
	Des *Players.Team
	WHETHER_ME  bool
}

func (e *JoinTeamEvent) Registry(curUser *Players.User, des *Players.Team) {
	e.currentUser = curUser
	e.Des = des
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<team:%s>", des.Name)
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}

func (e *JoinTeamEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message
}
