package events

import (
	"fmt"
	"konn/ingame/Players"
	"reflect"
)

type ColorChangeEvent struct {
	Event
	currentUser *Players.User
	before int
	after int
	WHETHER_ME  bool
}

func NewColorChangeEvent() {
	
}

func (e *ColorChangeEvent) Registry(curUser *Players.User, colorChange ...int) {
	e.currentUser = curUser
	e.before = colorChange[0]
	e.after = colorChange[1]
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<before:%d><after:%d>", colorChange[0], colorChange[1])
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
	EventSend(e.Event)
}

func (e *ColorChangeEvent) Receive(event *Event) {
	
}

func (e *ColorChangeEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message + br
}
