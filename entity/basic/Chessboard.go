package basic

const (
	MAPVERIFY = "<MAPNAME>"
)

type Chessboard struct {
	name  string
	grids [][]*Node
	size  int
}

func (ch *Chessboard) Name() string {
	return ch.name
}

func (ch *Chessboard) Grids() [][]*Node {
	return ch.grids
}

func (ch *Chessboard) Size() int {
	return ch.size
}
