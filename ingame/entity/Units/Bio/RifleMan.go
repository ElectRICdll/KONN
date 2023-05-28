package Bio

import (
	. "konn/ingame/entity/BasicClass"
)

type RifleMan struct {
	Creature
	Arming
}

func (u *RifleMan) String() string {
	return ""
}

func (u *RifleMan) initSub(userName string) {
	u.Man.initSub(userName)
	u.Health = 100
	u.Armor = 0
	u.Flexiblity = 1
	u.Arming = Arming{
		Damage: 5,
	}
}

func (u *RifleMan) Attack(oppo Substantive) {

}