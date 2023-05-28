package BasicClass

/*
Substance is a base class which defined how entities work in the game.
Almost every object will be Substance in a match.
*/
type Substance struct {
	id         int
	Health     int
	Armor      int
	Flexiblity int
	BelongsTo  string
}

func IDGenerate() int {
	return 0
}

func (u *Substance) initSub() {

}

func (u *Substance) ID() int {
	return u.id
}

func (u *Substance) Vanished() {

}