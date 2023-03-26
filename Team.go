package main

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
	name   string
	member map[string]*User
}

// TODO: map initialization
func NewTeam(teamName string) *Team {
	t := &Team{teamName, nil}
	return t
}

func (t *Team) String() string {
	members := func() string {
		s := ""
		for key, _ := range t.member {
			s += key + "\n"
		}
		return s
	}
	return "Team info: \nName: " + t.name + "\nMembers: \n" + members()
}

func (t *Team) AddMember(member string, collection map[string]*User) {
	t.member[member] = collection[member]
	// Temporal.
}

func (t *Team) RemoveMember(member string) {
	delete(t.member, member)
}
