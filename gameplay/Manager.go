package gameplay

import (
	"konn/entity/basic"
	"konn/utils"
)

var manager *Manager

type Manager struct {
	curRound  basic.Rounds
	totQueues []*utils.Queue
	players   []*basic.Player
}

func InstanceOfManager() *Manager {
	manager.Initialize()
	return manager
}

func (m *Manager) Initialize() {
	if m == nil {
		m = &Manager{
			curRound:  0,
			totQueues: make([]*utils.Queue, 0),
			players:   append(m.players, &basic.World),
		}
	}
}

func (m *Manager) AddQueue(queue *utils.Queue) {
	m.totQueues = append(m.totQueues, queue)
}

func (m *Manager) RemoveQueue(index int) {

}
