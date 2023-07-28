package basic

import (
	"fmt"
	"konn/entity"
)

type (
	Player entity.Player

	SubstanceID string
)

/*
Substance is a base class which defined how entities work in the game.
Almost every object will be Substance in a match.

At the mean time, Substance implemented prop.Substantive by default.
*/
type Substance struct {
	id     string
	size   int
	props  Properties
	belong Player
}

// TODO: id generate
func (s *Substance) InitSub() {
	s = &Substance{
		id: fmt.Sprintf("%s", &s),
		props: Properties{
			Health: 1,
		},
		size: 1,
	}
}

func (s *Substance) Vanished() {

}

func (s *Substance) ID() string {
	return s.id
}

func (s *Substance) Belong() entity.Player {
	return (entity.Player)(s.belong)
}

func (s *Substance) Props() Properties {
	return s.props
}

func (s *Substance) SetProps(properties Properties) {
	s.props = properties
}
