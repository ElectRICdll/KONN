package Players

import (
	"bufio"
	"encoding/json"
	"os"
)

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

var currentUser *User

type User struct {
	Name      string
	Color     string
	resources int
	belongsTo string
}

func NewUser(name string, color string) *User {
	return &User{name, color, 0, ""}
}

func Login() User {
	var res []byte
	var userData User

	userFile, err := os.Open("user.json")
	if err != nil {
		panic(nil) // TODO:
	}
	defer userFile.Close()
	inputReader := bufio.NewReader(userFile)
	inputReader.Read(res)

	json.Unmarshal(res, &userData)
	return userData
}

func (u *User) getActor() string {
	return u.String()
}

func (u *User) String() string {
	return "User info: \n" + u.Name + "\n" + u.belongsTo
	// TODO
}

func (u *User) setColor(c string) {
	u.Color = c
}

func (u *User) setBelong(hisTeam string) {
	u.belongsTo = hisTeam
}

func (u *User) resInitialize() {
	u.resources = 30000
}
