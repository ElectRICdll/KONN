package game

import "fmt"

type FrontierBase struct {
	Substance
	Building
	Constructor
	CList map[string]*Unit
	Producer
	health    int
	armor     int
	belongsTo string
}

func (b *FrontierBase) initSub(userName string) {
	b.health = 1000
	b.belongsTo = userName
}

// TODO: more fix 
func (b *FrontierBase) construct(obj Substance) (Substance, error) {
	fmt.Println("Constructing, it will take times...\n")
	defer fmt.Println("Construction completed.")
	if _, ok := obj.(Unit); ok {
		obj, _ = NewUnit("nil") // error handler needed.
	} else if _, ok := obj.(Building); ok {
		obj, _ = NewBuilding("nil") // error handler needed.
	} else {
		// error generater needed.
	}
	return obj, nil
}

func (b *FrontierBase) produce() {
	fmt.Println("Training...")
}

func (b *FrontireBase) Vanished() {

}

type InfantryBase struct {
	Substance
	Producer
	
	health    int
	armor     int
	belongsTo int
}

func (b *InfantryBase) initSub(username string) {

}

func (b *InfantryBase) produce() {

}

func (b *InfantryBase) Vanished() {

}

type Radar struct {
	Substance
	health    int
	armor     int
	belongsTo string
}

func (b *Radar) initSub(username string) {

}

func (b *Radar) Vanished() {

}