package main

import (
	"fmt"
	"konn/pre"
	"konn/tel"
	"konn/utils"
)

func sendEvent(e tel.Event) {
	eb := tel.EncodeEvent(e)
	fmt.Println("Committing event.")
	utils.Commiter(eb)
}


func main() {
	// e := &tel.ServiceEvent{}
	// tel.NewEvent(e, "New User: electric\nMes: Hello world!\n<END>")
	// sendEvent(e)
	pre.Menu()
}