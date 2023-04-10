package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type LocalData struct {
	user map[string]*User
	team map[string]*Team
}

func initialize() {
	fmt.Println("Initializing...\n")
	defer fmt.Println("Complete.")
	bootBag := &LocalData{map[string]*User{}, map[string]*Team{}}

	//
	bootBag.user["electric"] = NewUser("electric", "red")
	//
}

func menu() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("waiting for initialization...")
	initialize()
	fmt.Println("Done.")

	time.Sleep(1 * 1e9)
	fmt.Println("KONN. V" + VERSION)
	fmt.Println("Menu")
	fmt.Println("1.Online")
	fmt.Println("2.Settings")
	fmt.Println("3.Exit")
	fmt.Print("Select: ")
	// @magical
	switch cho := func(buf *bufio.Reader) string {defer exec.Command("cls");  res, _ := buf.ReadString('\n'); return res[:1]} (buf); {
		case cho == "1":
			Login()
		case cho == "2":
			// Settings()
		case cho == "3":
			return
	}
	// TODO: menu build
}
