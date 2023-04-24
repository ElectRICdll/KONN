package events

import (
	"fmt"
	"konn/ingame/basic"
	"reflect"
)

type DiseaseEvent struct {
	Event
	Target basic.Substantive
}

func (e *DiseaseEvent) Registry(target basic.Substantive) {
	e.Target = target
	msg := fmt.Sprintf("<Substance:%s>", e.Target.ID())
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}

func (e *DiseaseEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message
}