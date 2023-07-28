package basic

import (
	"konn/entity/prop"
	"konn/events"
	"konn/gameplay"
)

type Arming struct {
	Name      string
	Damage    int
	AntiArmor int
	Accuracy  int
}

func (a *Arming) Attack(target *prop.Substantive) {
	events.Register(events.NewAttackEvent(a, target))
	gameplay.HitJudge(*a, *target)
	gameplay.DamageHandler(*a, *target)
}

// TODO
func (a *Arming) Disable() {

}

// TODO
func (a *Arming) Launch() {

}
