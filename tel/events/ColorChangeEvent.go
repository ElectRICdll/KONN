package events

import (
	"fmt"
	. "konn/constants"
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
	msg := fmt.Sprintf("<user:%s><before:%d><after:%d>", e.currentUser.Name, e.before, e.after)
	e.Event = *NewEvent(reflect.TypeOf(*e).String(), msg)
	EventSend(e.Event)
}

func (e *ColorChangeEvent) Receive(event Event) {
	e.Event = event
	fmt.Sscanf(e.Message, "<user:%s><before:%d><after:%d>", e.currentUser.Name, e.before, e.after)
	// TODO: increase robustness
	if res := Players.Seeker(e.currentUser); res == nil {
		LogGenerate(WARNING, fmt.Sprintf("Missing player \"%s\"", e.currentUser.Name))
	} else {
		refer := reflect.ValueOf(&res)
		e.currentUser = refer.Interface().(*Players.User)
	}
}

func (e *ColorChangeEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message + br
}
