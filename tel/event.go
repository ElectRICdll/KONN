package tel

import (
	"reflect"
	"time"
)

type Event interface {
	Registry(msg string)
	String() string
}

func NewEvent(e Event, msg string) {
	e.Registry(msg)
	EncodeEvent(e)
}


type basicEvent struct {
	Event
	eventMsg string
}

func (e *basicEvent) Registry(msg string) {
	return 
}

type ServiceEvent struct {
	eventName string
	eventMsg  string
}

func (e *ServiceEvent) Registry(msg string) {
	e.eventName = reflect.TypeOf(e).String()
	e.eventMsg = msg
}

func (e *ServiceEvent) String() string {
	return "<BEGIN>EventName: " + e.eventName + ";EventDetails: " + e.eventMsg
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

type AttackEvent struct {
	*InGameEvent
}

type ContactEvent struct {
	*InGameEvent
}

type ConstructEvent struct {
	*InGameEvent
}

type DiseaseEvent struct {
	*InGameEvent
}

type ExposeEvent struct {
	*InGameEvent
}

type MovementEvent struct {
	*InGameEvent
}

type SetStatusEvent struct {
	*InGameEvent
}

type SetWeatherEvent struct {
	*InGameEvent
} // Beta.

type UnitAddEvent struct {
	*InGameEvent
}

type ChatEvent struct {
	Event
	eventName string
	eventMsg  string
	moment    time.Time
}

func (e *ChatEvent) registry(msg string) {
	e.eventName = reflect.TypeOf(e).String()
	e.eventMsg = msg
	e.moment = time.Now() // TODO: unpack time
}

type GlobalChatEvent struct {
	*ChatEvent
}

type PersonalChatEvent struct {
	*ChatEvent
}

type TeamInvitationEvent struct {
	*ChatEvent
}

type TeamChatEvent struct {
	*ChatEvent
} 






