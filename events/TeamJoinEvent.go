package events

import (
	"konn/entity"
)

type TeamJoinEvent struct {
	which *entity.User
	to    *entity.Team
}

func (e TeamJoinEvent) NewJoinTeamEvent(which *entity.User, to *entity.Team) Event {
	return TeamJoinEvent{
		which: which,
		to:    to,
	}
}

func (e TeamJoinEvent) toMessage() string {
	return e.String()
}

func (e TeamJoinEvent) String() string {
	return BEGIN + br
}
