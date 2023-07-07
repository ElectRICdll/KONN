package events

import (
	"konn/entity"
	"konn/entity/compaign"
)

type MoveEvent struct {
	Event
	which *entity.User
	from  *compaign.Position
	to    *compaign.Position
}

func NewMoveEvent(which *entity.User, from *compaign.Position, to *compaign.Position) Event {
	return MoveEvent{
		which: which,
		from:  from,
		to:    to,
	}
}

func (e MoveEvent) toMessage() string {
	return e.String()
}

func (e MoveEvent) String() string {
	return BEGIN + br
}
