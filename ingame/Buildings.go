package game

import "fmt"

type MainController struct {
	Building
	Constructor
	CList map[string]*Unit
	Producer
	health    int
	belongsTo string
}

func (b *MainController) initSub(userName string) {
	b.health = 1000
	b.belongsTo = userName
}

// TODO: more fix 
func (b *MainController) construct(obj Substance) (Substance, error) {
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

func (b *MainController) produce() {
	
}
