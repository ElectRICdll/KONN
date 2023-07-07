package events

import (
	"konn/entity"
)

type CreateTeamEvent struct {
	which *entity.User
	to    entity.Team
}

func NewCreateTeamEvent(which *entity.User, to entity.Team) Event {
	return CreateTeamEvent{
		which: which,
		to:    to,
	}
}

func (e CreateTeamEvent) toMessage() string {
	return e.String()
}

func (e CreateTeamEvent) String() string {
	return BEGIN + br + "EventName: "
}
