package entity

import (
	"errors"
	. "konn/ingame/basic"
	. "konn/ingame/mapping"
)

type Vehicle struct {
	Unit

	Substance
	VehicleSize int
}

func (u *Vehicle) initSub(userName string) {
	u.Substance = Substance{
		Health: 200,
		Armor: 100,
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

type Tank struct {
	Unit 
	
	Vehicle
	Gun Arming
	Turret Arming
}

func (u *Tank) initSub(userName string) {
	u.Vehicle.initSub(userName)
	u.Health = 300
	u.Armor = 200
	
	u.Gun = Arming{
		ItsName: "Gun",
		Damage: 50,
		AntiArmor: 300,
	}
	u.Turret = Arming{
		ItsName: "Turret",
		Damage: 20,
		AntiArmor: 1,
	}
}

func (u *Tank) Attack(oppo Substantive) {

}