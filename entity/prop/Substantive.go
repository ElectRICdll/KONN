package prop

import (
	"konn/entity/basic"
)

type Substantive interface {
	InitSub(belong *basic.Player)
	ID() string
	Vanished()
	Belong() basic.Player
	Props() basic.Properties
}
