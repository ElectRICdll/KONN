package bio

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type RifleMan struct {
	Man
	basic.Arming
}

func (u *RifleMan) String() string {
	return ""
}

func (u *RifleMan) initSub(userName string) {
	u.Man.InitSub(userName)
	u.Health = 100
	u.Armor = 0
	u.Flexiblity = 1
	u.Arming = basic.Arming{
		Damage: 5,
	}
}

func (u *RifleMan) Attack(oppo prop.Substantive) {

}
