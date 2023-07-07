package bio

import (
	"errors"
	"konn/entity/prop"

	"konn/entity/basic"
	. "konn/entity/compaign"
)

type Creature struct {
	prop.Unit
	basic.Substance
}

func (u *Creature) String() string {
	return ""
}

func (u *Creature) InitSub(userName string) {
	u.Substance = basic.Substance{
		Health:     50,
		Armor:      0,
		Flexiblity: 1,
		BelongsTo:  userName,
	}
}

func (u *Creature) Vanished() {

}

func (u *Creature) Movement(cur *Node, next *Node) error {
	var err error
	if IsConnected(cur, next) {
		err = nil

	} else {
		err = errors.New("Invalid movement!")
	}
	return err
}

// Structs under are inherited from struct Man.
