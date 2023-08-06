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
	IsDisable bool
	Belong    *Substance
}

func (a *Arming) Attack(target *prop.Substantive) {
	events.Register(events.NewAttackEvent(a, target))
	gameplay.Hit(*a, *target)
}

// TODO
func (a *Arming) Disable() {
	a.IsDisable = true
}

// TODO
func (a *Arming) Launch() {
	a.IsDisable = false
}
