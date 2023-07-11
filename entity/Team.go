package entity

import "strconv"

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
	id     int
	name   string
	member map[string]Player
}

// TODO: map initialization
func NewTeam(name string) *Team {
	return &Team{
		name:   name,
		member: make(map[string]Player),
	}
}

func (t *Team) String() string {
	members := func() string {
		s := ""
		for key, _ := range t.member {
			s += key + "\n"
		}
		return s
	}
	return "Team info: \nID: " + strconv.Itoa(t.id) + "\nName: " + t.name + "\nMembers: \n" + members()
}

//func (t *Team) AddMember(member string, collection map[string]Player) {
//	t.member[member] = collection[member]
//	// Temporal.
//}

func (t *Team) RemoveMember(member string) {
	delete(t.member, member)
}
