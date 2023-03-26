package main

/*
About a integration of User
Including:
	struct User include:
		string name(A user's name),
		string color(HISCOLOR)
		string belongsTo(The Team which this user belongs to)
	function NewUser
	method String, setColor, setBelong, resInitialize
*/

type User struct {
	name      string
	color     string
	resources int
	belongsTo string
}

func NewUser(name string, color string) *User {
	u := &User{name, color, 0, ""}
	return u
}

func (u *User) String() string {
	return "User info: \n" + u.name + "\n" + u.belongsTo
	// TODO
}

func (u *User) setColor(c string) {
	u.color = c
}

func (u *User) setBelong(hisTeam string) {
	u.belongsTo = hisTeam
}

func (u *User) resInitialize() {
	u.resources = 30000
}
