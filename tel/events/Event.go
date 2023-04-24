package events

const (
	newUserEvent  string = "NewUserEvent"
	userJoinEvent string = "UserJoinEvent"
	userLeftEvent string = "UserLeftEvent"
	userExitEvent string = "UserExitEvent"
	newTeamEvent  string = "NewTeamEvent"
	joinTeamEvent string = "JoinTeamEvent"
	leftTeamEvent string = "LeftTeamEvent"
)

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
	Type    string
	Message string
}

func NewEvent(eventName string, eventType string, msg string) *Event {
	return &Event{eventName, eventType, msg}
}

// func EventDecode(message string) {

// }
