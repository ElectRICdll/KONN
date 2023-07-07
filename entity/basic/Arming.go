package basic

type Arming struct {
	name      string
	damage    int
	antiArmor int
}

func (a *Arming) Attack() {
	//events.Register(events.NewAttackEvent())
}

func (a *Arming) Disable() {

}

func (a *Arming) Launch() {

}

func (a *Arming) Cancel() {

}

func (a *Arming) Name() string {
	return a.name
}

func (a *Arming) Damage() int {
	return a.damage
}

func (a *Arming) AntiArmor() int {
	return a.antiArmor
}
