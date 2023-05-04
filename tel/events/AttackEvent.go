package events

import (
	"fmt"
	. "konn/constants"
	"konn/ingame/Players"
	"konn/ingame/basic"
	"reflect"
)

type AttackEvent struct {
	Event
	currentUser *Players.User
	Attacker basic.Armer
	Target basic.Substantive
	ResDamage int
	WHETHER_ME bool
}

func (e *AttackEvent) Registry(curUser *Players.User, attacker basic.Armer, target basic.Substantive, damage int) {
	e.currentUser = curUser
	e.Attacker = attacker
	e.Target = target
	e.ResDamage = damage
	e.WHETHER_ME = true
	msg := fmt.Sprintf("<user:%s><attacker:%s><target:%s><damage:%d>", e.currentUser.Name, e.Attacker.ID(), e.Target.ID(), e.ResDamage)
	// should be e.Substantive.ID
	e.Event = *NewEvent(reflect.TypeOf(e).String(), msg)
}

func (e *AttackEvent) Receive(event Event) {
	e.Event = event
	fmt.Sscanf(e.Message, "<user:%s><attacker:%s><target:%s><damage:%d>", e.currentUser.Name, e.Attacker, e.Target, e.ResDamage)
	if res := Players.Seeker(e.currentUser); res == nil {
		LogGenerate(WARNING, fmt.Sprintf("Missing player \"%s\"", e.currentUser.Name))
	} else {
		// refer := reflect.ValueOf(&NewAttackEvent{})
		// e.Attacker = refer.Interface().(basic.Armer)
		// e.Target = 
	}
}

func (e *AttackEvent) String() string {
	return BEGIN + br + "EventName: " + e.Name + br + e.Event.Message + br
}