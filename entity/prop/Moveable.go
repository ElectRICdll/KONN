package prop

import (
	"konn/entity/basic"
	"konn/events"
)

type Moveable interface {
	MovementStart()
	MovementEnd()
}

func Movement(sub Moveable, from *basic.Node, to *basic.Node) {
	events.Register(events.NewMoveEvent(sub, from, to))
	sub.MovementStart()
	defer sub.MovementEnd()

}
