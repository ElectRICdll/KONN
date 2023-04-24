package main

import (
	"konn/tel"
	"konn/tui"
	"konn/utils"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-test" {
		var e tel.Event = &tel.InGameEvent{}
		tel.NewEvent(e, tel.AttackEvent, "User: electric\nMes: AttackEvent at tangyuanzi! {Damage: 30}\n<END>")
		utils.SendEvent(e)
	} else {
		utils.SetTerminalTitle("Meow!")
		Initialize()
		time.Sleep(2 * 1e9)
		ui.Init()
		tui.Menu()
	}
}
