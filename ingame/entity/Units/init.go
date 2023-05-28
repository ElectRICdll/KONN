package Units

import (
	. "konn/ingame/entity/BasicClass"
)

type Unit interface {
	Substantive
}

func NewUnit(it Unit) {
	it.initSub()
}
