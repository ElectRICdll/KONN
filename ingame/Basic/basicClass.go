package basic

type Substantive interface {
	initSub()
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
	Health    int
	Armor     int
	BelongsTo string
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