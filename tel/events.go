package tel

import (
	"reflect"
	"time"
)

const (
	NewUserEvent  string = "NewUserEvent"
	UserJoinEvent string = "UserJoinEvent"
	UserLeftEvent string = "UserLeftEvent"
	UserExitEvent string = "UserExitEvent"
	NewTeamEvent  string = "NewTeamEvent"
	JoinTeamEvent string = "JoinTeamEvent"
	LeftTeamEvent string = "LeftTeamEvent"
	TeamExitEvent string = "TeamExitEvent"
)

const (
	AttackEvent     string = "AttackEvent"
	ContactEvent    string = "ContactEvent"
	ConstructEvent  string = "ConstructEvent"
	DiseaseEvent    string = "DiseaseEvent"
	ExposeEvent     string = "ExposeEvent"
	MovementEvent   string = "MovementEvent"
	SetStatusEvent  string = "SetStatusEvent"
	SetWeatherEvent string = "SetWeatherEvent"
	UnitAddEvent    string = "UnitAddEvent"
)

const (
	GlobalChatEvent     string = "GlobalChatEvent"
	PersonalChatEvent   string = "PersonalChatEvent"
	TeamInvitationEvent string = "TeamInvitationEvent"
	TeamChatEvent       string = "TeamChatEvent"
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
	eventName string
	eventType reflect.Type
	eventMsg  string
}

func (e *basicEvent) Registry(msg string) {
	e.eventType = reflect.TypeOf(e.eventType)
	e.eventMsg = msg
}

type ServiceEvent struct {
	*basicEvent
}

func (e *ServiceEvent) Registry(eventName string, msg string) {
	e.basicEvent = &basicEvent{
		eventName: eventName,
		eventType: reflect.TypeOf(e),
		eventMsg:  msg,
	}
}

func (e *ServiceEvent) String() string {
	return "<BEGIN>EventName: " + e.eventType.String() + ";EventDetails: " + e.eventMsg
}

type InGameEvent struct {
	*basicEvent
}

type ChatEvent struct {
	*basicEvent
	moment time.Time
}

func (e *ChatEvent) registry(msg string) {
	e.eventType = reflect.TypeOf(e)
	e.eventMsg = msg
	e.moment = time.Now() // TODO: unpack time
}
