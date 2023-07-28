package gameplay

import (
	"bufio"
	"io"
	"konn/constants"
	"konn/entity/basic"
	"math"
	"os"
)

const (
	MAPPATH = ""
)

func MappingStart() {

}

func GetMapList() []basic.Chessboard {
	//mapFile, err = os.Open(MAPPATH)
	//if err != nil {
	//constants.LogGenerate()
	//}
	return nil
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
	if math.Abs(float64(n1.PosY()-n2.PosY())) <= 1 || math.Abs(float64(n1.PosY()-n2.PosY())) <= 1 {
		return true
	}
	return false
}
