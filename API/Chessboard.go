package api

import (
	"konn/entity/basic"
	"konn/gameplay"
)

func GetMapList() []string {
	res := make([]string, 0)
	for _, value := range gameplay.GetMapList() {
		res = append(res, value.String())
	}
	return res
}

func InitMap() string {

}

func AddMap(chessboard basic.Chessboard) {

}

func RemoveMap(chessboard basic.Chessboard) {

}
