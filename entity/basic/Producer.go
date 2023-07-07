package basic

import "konn/entity/prop"

type Producer struct {
	list map[string]*prop.Unit
}

func (p *Producer) Produce() {

}

func (p *Producer) List() map[string]*prop.Unit {
	return p.list
}
