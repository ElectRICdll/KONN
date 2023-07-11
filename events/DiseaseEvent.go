package events

import (
	"konn/entity/prop"
)

type DiseaseEvent struct {
	target prop.Substantive
}

func NewDiseaseEvent(target prop.Substantive) Event {
	return DiseaseEvent{
		target: target,
	}
}

func (e DiseaseEvent) ToMessage() string {
	return e.String()
}

func (e DiseaseEvent) String() string {
	return BEGIN + br + "EventName: "
}
