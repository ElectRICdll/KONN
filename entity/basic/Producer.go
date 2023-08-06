package basic

import (
	"konn/utils"
)

type Producer struct {
	IsDisable bool
	Belong    *Substance

	list  []Slot
	queue utils.Queue
}

func (p *Producer) Produce(index int) {
	utils.Logger.Info("Working on production...")
	defer utils.Logger.Info("Production completed.")
	p.queue.Push(p.list[index].work)
}

func (p *Producer) ShowList() []string {
	var result []string
	for _, value := range p.list {
		result = append(result, value.Name())
	}
	return result
}
