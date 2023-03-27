package main

import (
	"fmt"
	"time"
)

type Building() interface {
	isBuilding()
}

type Unit() interface {
	isUnit()
}

type Constructor interface {
	construct()
}

type Producer interface {
	product()
}

type Armer interface {
	attack()
}

// The things above is about the properties of a building.
type TeamsBelong interface {
	whichBelong() string
}

type Moveable interface {
	movement()
}

type Substance interface {
	setBody()
}

func NewBuilding(string itsName) Building {

}

func NewUnit(string itsName) Unit {

}
