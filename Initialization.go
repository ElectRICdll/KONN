package main

type DataBag {
	userBag map[string]*User
	teamBag map[string]*Team
}

func initialize() {
