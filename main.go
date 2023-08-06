package main

import (
	"os"
)

func main() {
	//tui.Initialize()
	if len(os.Args) > 1 && os.Args[1] == "-test" {
		APITest(os.Args[1:])
	} else {
		APIRun(os.Args)
	}
}
