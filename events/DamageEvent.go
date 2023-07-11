package events

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type DamageEvent struct {
	from *basic.Arming
	to   *prop.Substantive
	res  int
}

func NewDamageEvent(attacker *basic.Arming, target *prop.Substantive, damage int) Event {
	return DamageEvent{
		from: attacker,
		to:   target,
		res:  0,
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

func (e DamageEvent) ToMessage() string {
	return e.String()
}

func (e DamageEvent) String() string {
	return BEGIN + br + "EventName: "
}
