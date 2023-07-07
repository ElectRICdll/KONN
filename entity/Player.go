package entity

type Player struct {
	name string
	id   int
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) ID() int {
	return p.id
}
