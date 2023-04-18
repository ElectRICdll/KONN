package game

type Substance interface {
	initSub()
	Vanished()
}

type Constructor interface {
	Construct()
}

type Producer interface {
	Produce()
}

type Armer interface {
	Attack()
}

type TeamsBelong interface {
	WhichBelong() string
}

type Moveable interface {
	Movement()
}

type Building interface {
	Substance
}

type Unit interface {
	Substance
}

func NewBuilding(itsName string, pos Position) (Building, error) {
	var b Building
	return b, nil
}

func NewUnit(itsName string, pos Position) (Unit, error) {
	var u Unit
	return u, nil
}
