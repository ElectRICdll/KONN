package basic

import (
	"konn/constants"
)

type Node struct {
	id         int
	contains   [constants.BLOCK_SIZE]int
	curTerrain Terrain
	curWeather Weather
}

func NewNode(curTerrain Terrain) *Node {
	return &Node{
		curTerrain: curTerrain,
		curWeather: NewCommonWeather(),
	}
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
