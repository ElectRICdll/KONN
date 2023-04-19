package main

import (
	"fmt"
	"konn/pre"
	"konn/tel"
	"konn/utils"
	"os"
)

func sendEvent(e tel.Event) {
	eb := tel.EncodeEvent(e)
	fmt.Println("Committing event.")
	utils.Commiter(eb)
}


func main() { 
	if len(os.Args) > 1 && os.Args[1] == "-test" {
		e := &tel.NewUserEvent{}
		tel.NewEvent(e, "New User: electric\nMes: This is a new User!\n<END>")
		sendEvent(e)
	} else {
		pre.Menu()
	}
}