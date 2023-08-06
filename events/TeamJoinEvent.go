package events

import (
	"konn/entity/basic"
)

type TeamJoinEvent struct {
	which *basic.Player
	to    *basic.Team
}

func (e TeamJoinEvent) NewJoinTeamEvent(which *basic.Player, to *basic.Team) Event {
	return TeamJoinEvent{
		which: which,
		to:    to,
	}
}

func (e TeamJoinEvent) ToMessage() string {
	return e.String()
}

func (e TeamJoinEvent) String() string {
	return BEGIN + br
}
