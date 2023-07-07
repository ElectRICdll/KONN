package prop

type Substantive interface {
	InitSub()
	ID() int
	// Name()
	Vanished()
	BelongsTo() string
}
