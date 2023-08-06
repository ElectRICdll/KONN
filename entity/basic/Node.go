package basic

import "konn/utils"

const BLOCK_SIZE = 9

type Node struct {
	id         int
	contains   []*Substance
	remaining  int
	curTerrain Terrain
	curWeather Weather
	position   struct {
		Y int
		X int
	}
}

func NewNode(curTerrain Terrain) Node {
	return Node{
		curTerrain: curTerrain,
		curWeather: NewCommonWeather(),
	}
}

func (n *Node) ShowContains() []string {
	var result []string
	for _, element := range n.contains {
		result = append(result, element.Name())
	}
	return result
}

func (n *Node) Terrain() Terrain {
	return n.curTerrain
}

func (n *Node) SetTerrain(t Terrain) {
	n.curTerrain = t
}

func (n *Node) Weather() Weather {
	return n.curWeather
}

func (n *Node) SetWeather(w Weather) {
	n.curWeather = w
}

func (n *Node) PosY() int {
	return n.position.Y
}

func (n *Node) PosX() int {
	return n.position.X
}

func (n *Node) Enter(sub *Substance) {
	if n.remaining < sub.size {
		utils.Logger.Info("No more space to enter.")
	} else {
		n.contains = append(n.contains, sub)
		n.remaining -= sub.size
	}
}
