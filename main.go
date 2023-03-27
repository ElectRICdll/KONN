package main

import "fmt"

func main() {
	u := NewUser("electric", "red")

	fmt.Printf("Welcome! %s\n", u.name)
	operation()
}

func operation(u *User) {
	var choice int
	fmt.Println("Choose your next operation: ")
	fmt.Print("{1.Construction}\t{2.Attack}\t{3.Join a Team}")
	fmt.Scanf("%d", choice)
	switch choice {
		case 1:

	}
}
