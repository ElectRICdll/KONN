package prop

import (
	"konn/entity"
	"konn/entity/basic"
)

type Substantive interface {
	InitSub()
	ID() basic.ItemID
	Vanished()
	Belong() entity.Player
}
