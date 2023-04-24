package events

import (
	"fmt"
	. "konn/constants"
	"konn/ingame/Players"
	"reflect"
)

type ColorChangeEvent struct {
	*Event
	currentUser *Players.User
	before int
	after int
	WHETHER_ME  bool
}

func (e *ColorChangeEvent) Registry(curUser *Players.User, colorChange ...int) {
	e.before = colorChange[0]
	e.after = colorChange[1]
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<before:%s, after:%s>", reflect.TypeOf(e).String, colorChange[0], colorChange[1])
	NewEvent(e.Event, curUser.Name, msg)
}

func (e *ColorChangeEvent) Receive(event *Event) {

}

func (e *ColorChangeEvent) String() string {
	return BEGIN + br + "EventName: " + reflect.TypeOf(e).String() + br + e.Event.Message
}
