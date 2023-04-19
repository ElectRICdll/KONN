package units

import (
	"errors"

	. "konn/ingame/basic"
	. "konn/ingame/mapping"
)

type Man struct {
	Unit

	Substance
}

func (u *Man) String() string {
	return ""
}

func (u *Man) initSub(userName string) {
	u.Substance = Substance{
		Health:    50,
		Armor:     0,
		BelongsTo: userName,
	}
}

func (u *Man) Vanished() {

}

func (u *Man) Movement(cur *Node, next *Node) error {
	var err error
	if IsConnected(cur, next) {
		err = nil
		
	} else {
		err = errors.New("Invalid movement!")
	}
	return err
}

// Structs under are inherited from struct Man.
type RifleMan struct {
	Unit

	Man
	Arming
}

func (u *RifleMan) String() string {
	return ""
}

func (u *RifleMan) initSub(userName string) {
	u.Man.initSub(userName)
	u.Health = 100
	u.Armor = 0
	u.Arming = Arming{
		Damage: 5,
	}
}

func (u *RifleMan) Attack(oppo Substantive) {

}

