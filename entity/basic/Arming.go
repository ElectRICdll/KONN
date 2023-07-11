package basic

import (
	"konn/entity/prop"
	"konn/events"
)

type Arming struct {
	Name      string
	Damage    int
	AntiArmor int
	Accuracy  int
}

func (a *Arming) Attack(target *prop.Substantive) {
	events.Register(events.NewAttackEvent(a, target))
}

func (a *Arming) Disable() {

}

func (a *Arming) Launch() {

}
