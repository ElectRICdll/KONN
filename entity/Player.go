package entity

type Player struct {
	name   string
	id     int
	color  int
	belong *Team
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) ID() int {
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
