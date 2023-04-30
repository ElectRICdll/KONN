package events

import (
	"fmt"
	"konn/ingame/Players"
	"konn/ingame/mapping"
	"reflect"
)

type MoveEvent struct {
	Event
	currentUser *Players.User
	currentPos  *mapping.Position
	desination  *mapping.Position
	WHETHER_ME  bool
}

func (e *MoveEvent) Registry(curUser *Players.User, currentPos *mapping.Position, desination *mapping.Position) {
	e.currentUser = curUser
	e.currentPos = currentPos
	e.desination = desination
	e.WHETHER_ME = true
	msg := fmt.Sprintf("")
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}


func (e *MoveEvent) Receive(event *Event) {

}

func (e *MoveEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Message + br
}