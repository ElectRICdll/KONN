package game

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

func (u *Man) String() string {
	return ""
}

func (u *Man) initSub(userName string) {
	u.health = 50
	u.damage = 5
	u.belongsTo = userName
}

func (u *Man) attack(oppo *Substance) {
	fmt.Printf("Damage: %d! Attack: from %s", u.damage, reflect.TypeOf(u))
}
