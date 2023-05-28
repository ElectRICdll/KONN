package Map

import (
	"konn/ingame/entity/Properties"
)

type node struct {
	no         int
	contains   map[Properties.Substantive]int
	curTerrain Terrain
	curWeather Weather
}

func Newnode(curTerrain Terrain) *node {
	return &node{
		curTerrain: curTerrain,
		curWeather: NewCommonWeather(),
	}
}

func IsConnected(n1 *node, n2 *node) bool {
	return false
}

type Position interface {
	Pos() Position
}



