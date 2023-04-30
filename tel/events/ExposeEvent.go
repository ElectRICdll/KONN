package events

import "konn/ingame/Players"

type ExposeEvent struct {
	Event
	currentUser *Players.User
	
}