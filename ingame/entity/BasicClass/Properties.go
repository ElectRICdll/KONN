package BasicClass

/*
The Properties Class provides a standard of entities' basical properties.
*/
type Properties struct {
	Health    int
	Armor     int
	Damage    int
	AntiArmor int
	Scout     int
	AntiScout int
}

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
	ID() int
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


func NewBuilding(it Building) {
	it.initSub()
}






