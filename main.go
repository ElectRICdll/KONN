package main

import (
	"konn/entity"
	"konn/tel/events"
	"konn/utils"
	"os"
	"time"
)

func main() {
	//tui.Initialize()
	if len(os.Args) > 1 && os.Args[1] == "-test" {
		e := &events.ColorChangeEvent{}
		e.Registry(entity.NewUser("ElectRIC_dll", "Red"), 0, 1)
	} else {
		utils.SetTerminalTitle("Meow!")
		time.Sleep(2 * 1e9)
		//ui.Init()
		//tui.Menu()
	}
}
