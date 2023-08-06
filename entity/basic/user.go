package basic

var Self User

type User struct {
	id   string
	name string
	role Player
}

func (u *User) Login(username string, password string) {

}

func (u *User) Register(username string, password string, confirmation string) {

}

func (u *User) Logout() {

}

func (u *User) NewPassword(newPassword string) {

}
