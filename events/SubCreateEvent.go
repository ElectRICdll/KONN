package events

import (
	"konn/entity"
	"konn/entity/basic"
)

type SubCreateEvent struct {
	from *entity.Player
	to   basic.SubstanceID
}

func NewSubCreateEvent(from *entity.Player, to int) Event {
	return SubCreateEvent{
		from: from,
		to:   (basic.SubstanceID)(to),
	}
}

func (e SubCreateEvent) ToMessage() string {
	return ""
}
