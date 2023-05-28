package entity

import "reflect"

/*
About a integration of Team
Including:
	struct Team include:
		string name(The Team's name),
		map[string, User(Object)] member
	function NewTeam
	method String, AddMember, RemoveMember
*/

type Team struct {
	Name   string
	Member map[string]*User
}

// TODO: map initialization
func NewTeam(teamName string) *Team {
	t := &Team{teamName, nil}
	return t
}

func (t Team) getActor() string {
	return t.String()
}

func (t Team) getActorName() string {
	return t.Name
}

func (t Team) getActorType() string {
	return reflect.TypeOf(t).String()
}

func (t *Team) String() string {
	members := func() string {
		s := ""
		for key, _ := range t.Member {
			s += key + "\n"
		}
		return s
	}
	return "Team info: \nName: " + t.Name + "\nMembers: \n" + members()
}

func (t *Team) AddMember(member string, collection map[string]*User) {
	t.Member[member] = collection[member]
	// Temporal.
}

func (t *Team) RemoveMember(member string) {
	delete(t.Member, member)
}
