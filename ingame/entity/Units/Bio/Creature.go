package Bio

import (
	"errors"

	. "konn/ingame/entity/BasicClass"
	. "konn/ingame/entity/Units"
)

type Creature struct {
	Unit
	Substance
}

func (u *Creature) String() string {
	return ""
}

func (u *Creature) initSub(userName string) {
	u.Substance = Substance{
		Health:    50,
		Armor:     0,
		Flexiblity: 1,
		BelongsTo: userName,
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

