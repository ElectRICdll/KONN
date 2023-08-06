package basic

import "fmt"

const (
	MAPVERIFY = "<MAPNAME>"
)

type Chessboard struct {
	name  string
	grids [][]*Node
	ysize int
	xsize int
}

func NewChessBoard(name string, ysize int, xsize int) Chessboard {
	return Chessboard{
		name: name,
		grids: func(y int, x int) [][]*Node {
			grids := make([][]*Node, ysize)
			for i := 0; i < ysize; i++ {
				grids[i] = make([]*Node, xsize)
			}
			return grids
		}(ysize, xsize),
		ysize: ysize,
		xsize: xsize,
	}
}

func (ch *Chessboard) Name() string {
	return ch.name
}

func (ch *Chessboard) Grids() [][]*Node {
	return ch.grids
}

func (ch *Chessboard) SizeY() int {
	return ch.ysize
}

func (ch *Chessboard) SizeX() int {
	return ch.xsize
}

func (ch *Chessboard) String() string {
	return fmt.Sprintf("KONN map: %s\n")
}
