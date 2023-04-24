package events

import (
	"fmt"
	. "konn/constants"
	"konn/ingame/Players"
	"reflect"
)

type NewTeamEvent struct {
	*Event
	NewTeam *Players.Team
	WHETHER_ME  bool
}

func (e *NewTeamEvent) Registry(newTeam *Players.Team) {
	e.NewTeam = newTeam
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<new_team:%s>", e.NewTeam.Name)
	NewEvent(e.Event, newTeam.Name, msg)
}

func (e *NewTeamEvent) String() string {
	return BEGIN + br + "EventName: " + reflect.TypeOf(e).String() + br + e.Event.Message
}
