package machinary

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type Tank struct {
	Vehicle
	gun    basic.Arming
	turret basic.Arming
}

func (u *Tank) InitSub(userName string) {
	u.Substance.SetProps(basic.Properties{})
	u.VehicleSize = 1
	u.gun = basic.Arming{
		Name:      "Gun",
		Damage:    50,
		AntiArmor: 300,
	}
	u.turret = basic.Arming{
		Name:      "Turret",
		Damage:    20,
		AntiArmor: 1,
	}
}

func (u *Tank) Attack(oppo prop.Substantive) {

}
