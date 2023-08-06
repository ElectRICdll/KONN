package bio

import (
	"konn/entity/basic"
	"konn/entity/prop"
	"konn/utils"
)

type Man struct {
	Creature
	prop.Moveable
}

func (m *Man) InitSub() {
	m.Substance.SetProps(basic.Properties{
		Health: 50,
		Scout:  10,
	})
	m.isAlive = true
}

func (m *Man) MovementStart() {
	utils.Logger.Info("Moving.")
}

func (m *Man) MovementEnd() {

}
