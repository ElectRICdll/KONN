package basic

import (
	"konn/utils"
)

type Constructor struct {
	IsDisable bool
	Belong    *Substance

	list  []Slot
	queue utils.Queue
}

//func NewConstructor(Belong *Substance, list []Slot) *Constructor {
//
//}

// TODO
func (b *Constructor) Build(index int, grid *Node) {
	utils.Logger.Info("Constructing...")
	defer utils.Logger.Info("Completed.")
	b.queue.Push(b.list[index].work)
}

func (b *Constructor) Cancel(index int) {
	b.queue.Remove(index)
}

func (b *Constructor) ShowList() []string {
	var result []string
	for _, value := range b.list {
		result = append(result, value.Name())
	}
	return result
}

func (b *Constructor) complete() {

}
