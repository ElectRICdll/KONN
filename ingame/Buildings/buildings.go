package buildings

import (
	"fmt"
	. "konn/ingame/basic"
)

type FrontierBase struct {
	Building

	Substance
	Constructor
	Producer
}

func (b *FrontierBase) initSub(userName string) {

}

func (b *FrontierBase) Vanished() {

}

// TODO: more fix
func (b *FrontierBase) Construct(obj Substance) {
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

type InfantryBase struct {
	Building

	Substance
	Producer
}

func (b *InfantryBase) initSub(username string) {

}

func (b *InfantryBase) Produce() {

}

func (b *InfantryBase) Vanished() {

}

type Radar struct {
	Building

	Substance
}

func (b *Radar) initSub(username string) {

}

func (b *Radar) Vanished() {

}
