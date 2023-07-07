package entity

import (
	"fmt"
	"konn/entity/basic"
	"konn/entity/prop"
)

type FrontierBase struct {
	prop.Structure
	basic.Substance
	engineer basic.Builder
	train    basic.Producer
}

func (b *FrontierBase) initSub(userName string) {

}

func (b *FrontierBase) Vanished() {

}

// TODO: more fix
func (b *FrontierBase) Construct(obj basic.Substance) {
	fmt.Println("Constructing, it will take times...\n")
	defer fmt.Println("Construction completed.")
	// if _, ok := obj.(Unit); ok {
	// 	obj, _ = NewUnit("nil") // error handler needed.
	// } else if _, ok := obj.(Building); ok {
	// 	obj, _ = NewBuilding("nil") // error handler needed.
	// } else {
	// 	// error generater needed.
	// }
}

func (b *FrontierBase) produce() {
	fmt.Println("Training...")
}
