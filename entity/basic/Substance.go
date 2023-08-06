package basic

import (
	"fmt"
	"konn/events"
)

/*
Substance is a base class which defined how entities work in the game.
Almost every object will be Substance in a match.

At the mean time, Substance implemented prop.Substantive by default.
*/
type Substance struct {
	substanceID int
	id          string
	name        string
	size        int
	props       Properties
	belong      *Player
	isAlive     bool
}

// TODO: id generate
func (s *Substance) InitSub(belong *Player) {
	s = &Substance{
		substanceID: 1,
		id:          fmt.Sprintf("%s", &s),
		size:        1,
		belong:      belong,
		props: Properties{
			Health: 1,
		},
	}
}

func (s *Substance) Vanished() {
	events.Register(events.NewDiseaseEvent(s))
	s.isAlive = false
}

func (s *Substance) ID() string {
	return s.id
}

func (s *Substance) Name() string {
	return s.name
}

func (s *Substance) Belong() Player {
	return *s.belong
}

func (s *Substance) Props() Properties {
	return s.props
}

func (s *Substance) SetProps(properties Properties) {
	s.props = properties
}
