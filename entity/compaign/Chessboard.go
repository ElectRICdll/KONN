package compaign

import (
	"bufio"
	"io"
	"konn/constants"
	"os"
)

const (
	MAPVERIFY = "<MAPNAME>"
)

type Chessboard struct {
	name  string
	grids [][]*node
	size  int
}

// func ReadMaps() []Chessboard {

// }

func SearchMap(name string) Chessboard {
	file, err := os.Open(name)
	if err != nil {
		constants.LogGenerate(err.Error(), constants.WARNING)
	}
	inputReader := bufio.NewReader(file)

	for {
		str, err := inputReader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			constants.LogGenerate(constants.DEBUG, str)
		}
	}

	return Chessboard{}
}

func (ch *Chessboard) Name() string {
	return ch.name
}

func (ch *Chessboard) Grids() [][]*node {
	return ch.grids
}

func (ch *Chessboard) Size() int {
	return ch.size
}
