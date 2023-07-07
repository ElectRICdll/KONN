package basic

/*
Substance is a base class which defined how entities work in the game.
Almost every object will be Substance in a match.
*/
type Substance struct {
	Properties
	id         int
	flexiblity int
	belongsTo  string
}

func (u *Substance) initSub() {

}

func (u *Substance) Vanished() {

}

func (u *Substance) ID() int {
	return u.id
}

func (u *Substance) BelongsTo() string {
	return u.belongsTo
}
