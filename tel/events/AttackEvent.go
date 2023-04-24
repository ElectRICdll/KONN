package events

import (
	"fmt"
	"konn/ingame/Players"
	"konn/ingame/basic"
	"reflect"
)

type AttackEvent struct {
	Event
	currentUser *Players.User
	Attacker basic.Substantive
	Target basic.Substantive
	ResDamage int
	WHETHER_ME bool
}

func (e *AttackEvent) Registry(curUser *Players.User, attacker basic.Substantive, target basic.Substantive, damage int) {
	e.currentUser = curUser
	e.Attacker = attacker
	e.Target = target
	e.ResDamage = damage
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<Attacker:%s><Target:%s><Damage:%d>", e.Attacker.ID(), e.Target.ID(), e.ResDamage)
	// should be e.Substantive.ID
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}

func (e *AttackEvent) Receive(event *Event) {

}

func (e *AttackEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message + br
}