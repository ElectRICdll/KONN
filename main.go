package main

import "fmt"

func main() {
	u := NewUser("electric", "red")

	fmt.Printf("Welcome! %s\n", u.name)
}
