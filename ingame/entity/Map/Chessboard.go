package Map

import (
	"bufio"
	"io"
	. "konn/constants"
	"os"
)

const (
	MAPVERIFY = "<MAPNAME>"
)

type Chessboard struct {
	Mame string
	Grid [][]*node
	Size int
}

func ReadMap(name string) {
	file, err := os.Open(name)
	if err != nil {
		LogGenerate(err.Error(), WARNING)
	}
	inputReader := bufio.NewReader(file)

	for {
		str, err := inputReader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			LogGenerate(DEBUG, str)
		}
	}
}