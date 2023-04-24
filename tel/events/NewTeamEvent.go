package events

import (
	"fmt"
	"konn/ingame/Players"
	"reflect"
)

type NewTeamEvent struct {
	Event
	NewTeam *Players.Team
	WHETHER_ME  bool
}

func (e *NewTeamEvent) Registry(newTeam *Players.Team) {
	e.NewTeam = newTeam
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<new_team:%s>", e.NewTeam.Name)
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}

func (e *NewTeamEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message
}
