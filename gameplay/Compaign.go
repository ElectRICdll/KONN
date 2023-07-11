package gameplay

import (
	"bufio"
	"io"
	"konn/constants"
	"konn/entity/basic"
	"os"
)

func MappingStart() {

}

func SearchMap() []basic.Chessboard {

}

func ReadMap(name string) basic.Chessboard {
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

	return basic.Chessboard{}
}

func IsConnected(n1 *basic.Node, n2 *basic.Node) bool {
	return false
}
