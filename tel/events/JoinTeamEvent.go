package events

import (
	"fmt"
	. "konn/constants"
	"konn/ingame/Players"
	"reflect"
)

type JoinTeamEvent struct {
	*Event
	currentUser *Players.User
	Des *Players.Team
	WHETHER_ME  bool
}

func (e *JoinTeamEvent) Registry(curUser *Players.User, des *Players.Team) {
	e.currentUser = curUser
	e.Des = des
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<team:%s>", des.Name)
	NewEvent(e.Event, curUser.Name, msg)
}

func (e *JoinTeamEvent) String() string {
	return BEGIN + br + "EventName: " + reflect.TypeOf(e).String() + br + e.Event.Message
}
