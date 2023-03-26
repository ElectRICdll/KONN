package main

import (
	"fmt"
	"time"
)

type Construction interface {
	constructSelf()
	ruinSelf()
}

type Production interface {
	product()
}

type Arm interface {
	attack()
}

// The things above is about the properties of a building.

type Basical_Building struct {
	health        int
	constructTime time.Duration
	C             bool
	P             bool
	A             bool
	belongTo      string
}

// Defined a basic building.

func NewBuilding() {
	b := &Basical_Building{300, 10, true, false, false, ""}
	return b
}

func (b *Basical_Building) constructSelf() {
	fmt.Printf("Constructing...\n")
	time.Sleep(b.constructTime * 1e9)
	fmt.Printf("Construction complete.\n")
}

func (b *Basical_Building) ruinSelf() {
	fmt.Printf("Lost a building.\n")
}
