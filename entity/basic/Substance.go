package basic

import "konn/entity"

type (
	Player entity.Player

	SubstanceID int
	ItemID      int
)

/*
Substance is a base class which defined how entities work in the game.
Almost every object will be Substance in a match.

At the mean time, Substance implemented prop.Substantive by default.
*/
type Substance struct {
	id     ItemID
	props  Properties
	belong Player
}

func (s *Substance) InitSub() {
	s = &Substance{
		props: Properties{
			Health: 1,
		},
	}
}

func (s *Substance) Vanished() {

}

func (s *Substance) ID() ItemID {
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
