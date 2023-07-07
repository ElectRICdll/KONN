package Machinary

type Tank struct {
	Unit

	Vehicle
	Gun    Arming
	Turret Arming
}

func (u *Tank) initSub(userName string) {
	u.Vehicle.initSub(userName)
	u.Health = 300
	u.Armor = 200

	u.Gun = Arming{
		ItsName:   "Gun",
		Damage:    50,
		AntiArmor: 300,
	}
	u.Turret = Arming{
		ItsName:   "Turret",
		Damage:    20,
		AntiArmor: 1,
	}
}

func (u *Tank) Attack(oppo Substantive) {

}