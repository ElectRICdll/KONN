package main

import (
	"konn/ingame/Players"
	"konn/tel/events"
	"konn/tui"
	"konn/utils"
	"os"
	"time"

	ui "github.com/gizak/termui/v3"
)

func main() {
	tui.Initialize()
	if len(os.Args) > 1 && os.Args[1] == "-test" {
		e := &events.ColorChangeEvent{}
		e.Registry(Players.NewUser("ElectRIC_dll", "Red"), 0, 1)
	} else {
		utils.SetTerminalTitle("Meow!")
		time.Sleep(2 * 1e9)
		ui.Init()
		tui.Menu()
	}
}
