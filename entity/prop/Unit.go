package prop

type Unit interface {
	Substantive
}

func NewUnit(it Unit) {
	it.InitSub()
}
