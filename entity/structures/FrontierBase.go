package entity

import (
	"konn/entity/basic"
	"konn/entity/prop"
)

type FrontierBase struct {
	prop.Structure
	basic.Substance
	defence        basic.Arming
	engineer       basic.Constructor
	trainingCourse basic.Producer
}

// TODO
func NewFrontierBase(belong *basic.Player) *FrontierBase {
	return &FrontierBase{
		defence: basic.Arming{
			Name:      "",
			Damage:    20,
			AntiArmor: 10,
			Accuracy:  85,
			IsDisable: true,
		},
		engineer: basic.Constructor{
			IsDisable: true,
		},
		trainingCourse: basic.Producer{
			IsDisable: true,
		},
	}
}

func (b *FrontierBase) InitSub(userName string) {
	b.SetProps(basic.Properties{
		Health:    400,
		Armor:     50,
		Scout:     40,
		AntiScout: 0,
		Flex:      0,
	})
}

func (b *FrontierBase) Vanished() {
	b.engineer.IsDisable = true
	b.trainingCourse.IsDisable = true
}
