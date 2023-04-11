package pre

import (
	"bufio"
	"fmt"
	"konn/info"
	game "konn/ingame"
	"os"
	"os/exec"
	"time"
)

type LocalData struct {
	user map[string]*game.User
	team map[string]*game.Team
}

func initialize() {
	fmt.Println("Initializing...\n")
	defer fmt.Println("Complete.")
	bootBag := &LocalData{map[string]*game.User{}, map[string]*game.Team{}}

	//
	bootBag.user["electric"] = game.NewUser("electric", "red")
	//
}

func Menu() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("waiting for initialization...")
	initialize()
	fmt.Println("Done.")

	time.Sleep(1 * 1e9)
	fmt.Println("KONN. V" + info.VERSION)
	fmt.Println("Menu")
	fmt.Println("1.Online")
	fmt.Println("2.Settings")
	fmt.Println("3.Exit")
	fmt.Print("Select: ")
	// @magical
	switch cho := func(buf *bufio.Reader) string {defer exec.Command("cls");  res, _ := buf.ReadString('\n'); return res[:1]} (buf); {
		case cho == "1":
			game.Login()
		case cho == "2":
			// Settings()
		case cho == "3":
			return
	}
	// TODO: menu build
}

func ChatSys() {
	
}
