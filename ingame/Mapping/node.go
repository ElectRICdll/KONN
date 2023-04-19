package mapping

import (
	. "konn/ingame/basic"
)

type Position interface {
	Pos() Position
}

type Node struct {
	no        int
	contains  map[Substantive]int
	neighbors []int
}

func IsConnected(n1 *Node, n2 *Node) bool {
	return false
}

