package basic

import (
	"konn/entity/prop"
	"konn/utils"
)

type Producer struct {
	list  map[string]*prop.Unit
	queue utils.Queue
}

func (p *Producer) Produce(name string) {
	p.queue.Push(p.list[name])
}

func (p *Producer) List() map[string]*prop.Unit {
	return p.list
}
