package gameplay

import (
	"konn/entity/basic"
	"konn/utils"
	"math"
)

const (
	MAPPATH = ""
)

func MappingStart() {

}

func GetMapList() []basic.Chessboard {
	var res []basic.Chessboard = make([]basic.Chessboard, 0)
	files, err := utils.DirScanner(MAPPATH)
	if err != nil {
		utils.Logger.Error(err.Error())
	} else {
		for _, file := range files {
			res = append(res, readMap(file))
		}
	}
	return nil
}

func readMap(name string) basic.Chessboard {
	//file, err := os.Open(name)
	//if err != nil {
	//	utils.Logger.Error(err.Error())
	//}
	//inputReader := bufio.NewReader(file)

	// TODO: Map decode

	return basic.Chessboard{}
}

func IsConnected(n1 *basic.Node, n2 *basic.Node) bool {
	if math.Abs(float64(n1.PosY()-n2.PosY())) <= 1 || math.Abs(float64(n1.PosY()-n2.PosY())) <= 1 {
		return true
	}
	return false
}
