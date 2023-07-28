package events

import (
	"konn/entity/basic"
)

type SubCreateEvent struct {
	from *basic.Builder
	to   basic.SubstanceID
}

func NewSubCreateEvent(from *basic.Builder, to basic.SubstanceID) Event {
	return SubCreateEvent{
		from: from,
		to:   to,
	}
}

func (e SubCreateEvent) ToMessage() string {
	return ""
}
