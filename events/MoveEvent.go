package events

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type MoveEvent struct {
	Event
	which *prop.Substantive
	from  *basic.Node
	to    *basic.Node
}

func NewMoveEvent(which *prop.Substantive, from *basic.Node, to *basic.Node) Event {
	return MoveEvent{
		which: which,
		from:  from,
		to:    to,
	}
}

func (e MoveEvent) ToMessage() string {
	return e.String()
}

func (e MoveEvent) String() string {
	return BEGIN + br
}
