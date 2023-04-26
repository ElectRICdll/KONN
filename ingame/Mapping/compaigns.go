package mapping

import (
	"bufio"
	"io"
	. "konn/constants"
	"os"
)

const (
	MAPVERIFY = "<MAPNAME>"
)

// type ComType interface {
// 	decodeCom()
// 	encodeCom()
// }

type Compaign struct {
	name     string
	plot     *[][]Node
	plotSize int
}

func TestCompaign() {

} // test map

func ReadCompaign(name string) {
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
