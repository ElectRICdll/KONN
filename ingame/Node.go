package game

type Position interface {
	Pos() Position
}

type node struct {
	no        int
	contains  map[Substance]int
	neighbors []int
}