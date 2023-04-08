package main

import "fmt"

type MainController struct {
	Building
	Constructor
	CList map[string]*Unit
	Producer
	health    int
	belongsTo string
}

func (self *MainController) initSub(userName string) {
	self.health = 1000
	self.belongsTo = userName
}

// TODO: more fix
func (self *MainController) construct(obj Substance) (Substance, error) {
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

func (self *MainController) produce() {

}
