package events

import (
	"konn/entity"
)

type CreateTeamEvent struct {
	which *entity.Player
	to    entity.Team
}

func NewCreateTeamEvent(which *entity.Player, to entity.Team) Event {
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
