package basic

import (
	"konn/entity/prop"
	"konn/events"
)

type Builder struct {
	list map[string]prop.Structure
}

// TODO
func (b *Builder) Build(id SubstanceID, grid *Node) {
	events.Register(events.NewSubCreateEvent(b, id))

}

// TODO
func (b *Builder) Cancel() {

}

// TODO
func (b *Builder) List() map[string]prop.Structure {
	return b.list
}
