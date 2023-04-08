package main

import "fmt"

type DataBag struct {
	user map[string]*User
	team map[string]*Team
}

func initialize() {
	fmt.Println("Initializing...\n")
	defer fmt.Println("Complete.")
	bootBag := &DataBag{map[string]*User{}, map[string]*Team{}}

	//
	bootBag.user["electric"] = NewUser("electric", "red")
	//
}

func menu(u *User) {
	fmt.Println("KONN. V" + VERSION)
	fmt.Println("Menu")
	// TODO: menu build
}
