package bio

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type RifleMan struct {
	Man
	rifle basic.Arming
}

func (m *RifleMan) InitSub(userName string) {
	m.Substance = basic.Substance{}
	m.rifle = basic.Arming{}
}

func (m *RifleMan) Attack(oppo prop.Substantive) {
	//events.Register(events.NewAttackEvent())
}

func (m *RifleMan) Movement() {
	//events.NewMoveEvent()
}

func (m *RifleMan) String() string {
	return ""
}
