package machinary

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type Machine struct {
	prop.Unit
	basic.Substance
	isAlive bool
}

func (u *Machine) Vanished() {

}
