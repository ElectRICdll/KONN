package tui

import (
	"fmt"
	. "konn/constants"
	"konn/ingame/Players"
	"konn/utils"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// TODO: tui!!!

type LocalData struct {
	user map[string]*Players.User
	team map[string]*Players.Team
}

func Initialize() {
	fmt.Print("Initializing...")
	defer fmt.Println("Done.")
	bootBag := &LocalData{map[string]*Players.User{}, map[string]*Players.Team{}}
	toServer := utils.ClientGenerate()
	var RecvMessage string chan
	go utils.Receiver(toServer, RecvMessage)
	// undetermined code
	bootBag.user["electric"] = Players.NewUser("electric", "red")
	//
}

func Menu() {
	menu := utils.SetMainMenu(
		"KONN. Version:"+VERSION,
		[]string{"1.Online", "2.Library", "3.Settings", "4.Exit"},
	)
	menu.SelectedRowStyle = ui.NewStyle(ui.ColorYellow)

	subMenu := []*widgets.List{
		utils.SetSubMenu(
			"Online Competitions",
			[]string{"Connecting to online services..."}),
		utils.SetSubMenu(
			"Library",
			[]string{"Buildings", "Units", "Map", "Weather"}),
		utils.SetSubMenu("Settings",
			[]string{"Draw Frames", "Server IP:" + Host, "Server Port"}),
	}
	for i := 0; i < len(subMenu); i++ {
		subMenu[i].SelectedRowStyle = ui.NewStyle(ui.ColorYellow)
	}

	switcher := utils.SetSwitcher("Menu", "SubMenu")

	ui.Render(menu)
	curSub := utils.SetSubMenu("", nil)
	renderSwitch := func() {
		switch switcher.ActiveTabIndex {
		case 0:
			ui.Render(menu)
		case 1:
			ui.Render(menu, curSub)
		}
	}
	controlSwitch := func() *widgets.List {
		switch switcher.ActiveTabIndex {
		case 0:
			return menu
		case 1:
			return curSub
		default:
			return menu
		}
	}

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "s", "<Down>":
			control := controlSwitch()
			control.ScrollDown()
			// menu.ScrollDown() || curSub.ScrollDown()
		case "w", "<Up>":
			control := controlSwitch()
			control.ScrollUp()
		case "<Escape>":
			switcher.FocusLeft()
			ui.Clear()
		case "<Enter>":
			switcher.FocusRight()
			ui.Clear()
			if switcher.ActiveTabIndex == 0 {
				curSub = subMenu[menu.SelectedRow]
			} else if menu.SelectedRow != 2 {

			} else {

			}
		}
		renderSwitch()
	}
}

func ChatSys() {

}
