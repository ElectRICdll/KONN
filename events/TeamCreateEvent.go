package events

import (
	"konn/entity/basic"
)

type CreateTeamEvent struct {
	which *basic.Player
	to    basic.Team
}

func NewCreateTeamEvent(which *basic.Player, to basic.Team) Event {
	return CreateTeamEvent{
		which: which,
		to:    to,
	}
}

func (e CreateTeamEvent) ToMessage() string {
	return e.String()
}

func (e CreateTeamEvent) String() string {
	return BEGIN + br + "EventName: "
}
