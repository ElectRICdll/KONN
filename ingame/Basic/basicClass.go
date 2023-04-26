package basic

type Substantive interface {
	initSub()
	ID() int
	// Name()
	Vanished()
}

type ConstructSub interface {
	Construct()
}

type ProduceSub interface {
	Produce()
}

type Armer interface {
	Attack()
}

type TeamsBelong interface {
	WhichBelong() string
}

type Moveable interface {
	Movement()
}

type Building interface {
	Substantive
}

type Unit interface {
	Substantive
}

func NewBuilding(it Building) {
	it.initSub()
}

func NewUnit(it Unit) {
	it.initSub()
}

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

type Arming struct {
	ItsName   string
	Damage    int
	AntiArmor int
}

type Constructor struct {
	List map[string]*Building
}

type Producer struct {
	List map[string]*Unit
}
