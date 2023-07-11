package machinary

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type Vehicle struct {
	Machine
	prop.Moveable
	VehicleSize int
}

func (u *Vehicle) InitSub(userName string) {
	u.Substance.SetProps(basic.Properties{})
	u.VehicleSize = 1
	u.isAlive = true
}

func (u *Vehicle) Vanished() {
	u.isAlive = false
}

func (u *Vehicle) Movement() {
	//var err error
	//if IsConnected(cur, next) {
	//	err = nil
	//
	//} else {
	//	err = errors.New("Invalid movement!")
	//}
	//return err
}
