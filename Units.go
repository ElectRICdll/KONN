package main

import (
	"fmt"
	"reflect"
)

type Man struct {
	Unit
	// Moveable
	Armer
	damage    int
	health    int
	belongsTo string
}

func (self *Man) initSub(userName string) {
	self.health = 50
	self.damage = 5
	self.belongsTo = userName
}

func (self *Man) attack(oppo *Substance) {
	fmt.Printf("Damage: %d! Attack: from %s", self.damage, reflect.TypeOf(self))
}
