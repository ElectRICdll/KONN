package bio

import (
	"konn/entity/basic"
	"konn/entity/prop"
	"konn/events"
)

type Creature struct {
	prop.Unit
	basic.Substance
	isAlive bool
}

func (c *Creature) IsAlive() bool {
	return c.isAlive
}

func (c *Creature) Vanished() {
	events.Register(events.NewDiseaseEvent(c))
	c.isAlive = false
}
