package pre

import (
	"fmt"
	"konn/constants"
	game "konn/ingame"
	"konn/utils"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// TODO: tui!!!

type LocalData struct {
	user map[string]*game.User
	team map[string]*game.Team
}

func initialize() {
	fmt.Print("Initializing...")
	defer fmt.Println("Done.")
	bootBag := &LocalData{map[string]*game.User{}, map[string]*game.Team{}}
	
	// TODO: undetermined code
	bootBag.user["electric"] = game.NewUser("electric", "red")
	//
}

func Menu() {
	utils.SetTerminalTitle("Meow!")
	initialize()
	time.Sleep(2 * 1e9)
	ui.Init()
	
	menu := utils.SetMainMenu(
		"KONN. Version:" + constants.VERSION,
		[]string{"1.Online","2.Library","3.Settings","4.Exit",},
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
			[]string{"Draw Frames", "Server IP", "Server Port"}), 
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
		case "<Enter>":
			switcher.FocusRight()
			ui.Clear()
			curSub = subMenu[menu.SelectedRow]
			// renderSwitch()
		case "<Escape>":
			switcher.FocusLeft()
			ui.Clear()
		}
		renderSwitch()
	}
}

func ChatSys() {
	
}
