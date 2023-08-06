package basic

var World Player = Player{
	id:     "world",
	name:   "world",
	color:  0,
	belong: nil,
}

type Player struct {
	id     string
	name   string
	color  int
	belong *Team
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) ID() string {
	return p.id
}

// TODO: parse int to color's name.
func (p *Player) Color() string {
	return ""
}

func (p *Player) SetColor(c int) {
	p.color = c
}

func (p *Player) Belong() Team {
	return *p.belong
}

func (p *Player) SetBelong(id int) {

}
