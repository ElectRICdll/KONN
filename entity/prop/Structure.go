package prop

type Structure interface {
	Substantive
}

func NewStructure(it Structure) {
	it.InitSub()
}
