package entity

import (
	. "konn/entity/basic"
	. "konn/entity/prop"
)

type InfantryBase struct {
	Structure
	Substance
	Producer
}

func (b *InfantryBase) initSub(username string) {

}

func (b *InfantryBase) Produce() {

}

func (b *InfantryBase) Vanished() {

}
