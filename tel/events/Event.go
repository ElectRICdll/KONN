package events

import (
	"fmt"
	. "konn/constants"
	"konn/utils"
)

// const (
// 	newUserEvent  string = "NewUserEvent"
// 	userJoinEvent string = "UserJoinEvent"
// 	userLeftEvent string = "UserLeftEvent"
// 	userExitEvent string = "UserExitEvent"
// 	newTeamEvent  string = "NewTeamEvent"
// 	joinTeamEvent string = "JoinTeamEvent"
// 	leftTeamEvent string = "LeftTeamEvent"
// )

// const (
// 	AttackEvent     string = "AttackEvent"
// 	ContactEvent    string = "ContactEvent"
// 	ConstructEvent  string = "ConstructEven
// 	ConstructEvent  string = "ConstructEvent"
// 	DiseaseEvent    string = "DiseaseEvent"
// 	ExposeEvent     string = "ExposeEvent"
// 	MovementEvent   string = "MovementEvent"
// 	SetStatusEvent  string = "SetStatusEvent"
// 	SetWeatherEvent string = "SetWeatherEvent"
// 	UnitAddEvent    string = "UnitAddEvent"
// )

// const (
// 	GlobalChatEvent     string = "GlobalChatEvent"
// 	PersonalChatEvent   string = "PersonalChatEvent"
// 	TeamInvitationEvent string = "TeamInvitationEvent"
// 	TeamChatEvent       string = "TeamChatEvent"

type Event struct {
	Name    string
	Message string
}

func NewEvent(eventName string, msg string) *Event {
	return &Event{eventName, msg}
}

func EventSend(e Event) {
	utils.Commiter(
		fmt.Sprintf(EVENTBAG + MODULE_NAME + "\n" +
			THE_VERSION + VERSION + "\n" +
			EVENTTYPE + e.Name + "\n" +
			BODY + e.Message + BODYEND,
		),
	)
	LogGenerate(DEBUG, "Committing event...")
}

func EventRecv(msg string) {
	var (
		e Event
		module_Msg string
		version_Msg string
	)
	fmt.Sscanf(msg, 
		EVENTBAG + "%s" + "\n" +
		THE_VERSION + "%s" + "\n" +
		EVENTTYPE + "%s" + "\n" +
		BODY + "%s" + BODYEND,
		module_Msg, version_Msg, e.Name, e.Message,
	)
	switch e.Name {
	case "NewUserEvent":
		
	}
}

func (e *Event) String() string {
	return "EventName: " + e.Name + br + e.Message + br
}