package api

import (
	"konn/entity/compaign"
)

func GetMapList() []compaign.Chessboard {
	return compaign.ReadMaps()
}

func InitMap() {

}

func AddMap(chessboard compaign.Chessboard) {

}

func RemoveMap(chessboard compaign.Chessboard) {

}
