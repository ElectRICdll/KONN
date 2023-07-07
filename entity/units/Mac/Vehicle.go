package entity

import (
	"errors"
	. "konn/entity/basic"
	. "konn/entity/prop"
)

type Vehicle struct {
	Unit

	Substance
	VehicleSize int
}

func (u *Vehicle) initSub(userName string) {
	u.Substance = Substance{
		Health:    200,
		Armor:     100,
		BelongsTo: userName,
	}
	u.VehicleSize = 1
}

func (u *Vehicle) Vanished() {

}

func (u *Vehicle) Movement(cur *Node, next *Node) error {
	var err error
	if IsConnected(cur, next) {
		err = nil

	} else {
		err = errors.New("Invalid movement!")
	}
	return err
}
