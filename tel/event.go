package tel

import "time"

type Event interface {
	registry()
}

type basicEvent struct {
	Event
	eventMsg string
}

type ServiceEvent struct {
	Event
	eventName string
	eventMsg  string
}

func (e *ServiceEvent) registry() {

}

type NewUserEvent struct {
	*ServiceEvent
}

type UserJoinEvent struct {
	*ServiceEvent
}

type UserLeftEvent struct {
	*ServiceEvent
}

type UserExitEvent struct {
	*ServiceEvent
}

type NewTeamEvent struct {
	*ServiceEvent
}

type JoinTeamEvent struct {
	*ServiceEvent
}

type LeftTeamEvent struct {
	*ServiceEvent
}

type TeamExitEvent struct {
	*ServiceEvent
}

type InGameEvent struct {
	Event
	eventName string
	eventMsg  string
}

type ChatEvent struct {
	Event
	eventName string
	eventMsg  string
	moment    time.Time
}