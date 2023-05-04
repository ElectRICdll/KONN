package events

import (
	"konn/ingame/Players"
	"konn/ingame/basic"
)

const (
	TO_ALL    = "all"
	TO_MYTEAM = "myteam"
	TO_SINGLE = "singleton"
)

type EventBag interface {
	getEType() string
	Registry()
	Receive()
	Send() string
}

type ServiceEvent interface {
	Registry(*Players.Actor, ...int)
	Receive(e *Event)
	String() string
}

type GameEvent interface {
	Registry(...*basic.Substance)
	Receive(e *Event)
	String() string
}

type ChatEvent interface {
	Registry(*Players.Actor, string)
	Receive(e *Event)
	String() string
}