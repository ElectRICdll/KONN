package main

type Substance interface {
	initSub()
	vanished()
}

type Constructor interface {
	construct()
}

type Producer interface {
	produce()
}

type Armer interface {
	attack()
}

type TeamsBelong interface {
	whichBelong() string
}

type Moveable interface {
	move()
}

type Building interface {
	Substance
}

type Unit interface {
	Substance
}

func NewBuilding(itsName string) (Building, error) {
	var b Building
	return b, nil
}

func NewUnit(itsName string) (Unit, error) {
	var u Unit
	return u, nil
}
