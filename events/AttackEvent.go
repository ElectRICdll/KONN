package events

import (
	"konn/entity"
	"konn/entity/basic"
	"konn/entity/prop"
)

type AttackEvent struct {
	currentUser *entity.User
	attacker    *basic.Arming
	target      *prop.Substantive
	resDamage   int
	isMe        bool
}

func (e AttackEvent) NewAttackEvent(curUser *entity.User, attacker *basic.Arming, target *prop.Substantive, damage int) Event {
	return AttackEvent{
		currentUser: curUser,
		attacker:    attacker,
		target:      target,
		resDamage:   0,
		isMe:        true,
	}
}

//func (e *AttackEvent) Receive(event Event) {
//	e.Event = event
//	fmt.Sscanf(e.Message, "<user:%s><attacker:%s><target:%s><damage:%d>", e.currentUser.Name, e.Attacker, e.Target, e.ResDamage)
//	if res := entity.Player.Seeker(e.currentUser); res == nil {
//		LogGenerate(WARNING, fmt.Sprintf("Missing player \"%s\"", e.currentUser.Name))
//	} else {
//		// refer := reflect.ValueOf(&NewAttackEvent{})
//		// e.Attacker = refer.Interface().(basic.Armer)
//		// e.Target =
//	}
//}

func (e AttackEvent) toMessage() string {
	return e.String()
}

func (e AttackEvent) String() string {
	return BEGIN + br + "EventName: "
}
