package tel

import "konn/ingame/Players"

type ServiceEvents interface {
	Registry(*Players.Actor, ...int)
	Receive(e *Event)
	String() string
}

type GameEvent interface {
	Registry(...*ingame.Substance)
	Receive(e *Event)
	String() string
}

type ChatEvent interface {
	Registry(*Players.Actor, string)
	Receive(e *Event)
	String() string
}

const (
	TO_ALL    = "all"
	TO_MYTEAM = "myteam"
	TO_SINGLE = "singleton"
)