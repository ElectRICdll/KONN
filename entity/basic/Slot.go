package basic

type Rounds int

type Slot struct {
	work     func()
	name     string
	duration Rounds
	cost     int
}

func NewSlot(do func(), name string, duration int, cost int) Slot {
	return Slot{
		work:     do,
		name:     name,
		duration: (Rounds)(duration),
		cost:     cost,
	}
}

func (s Slot) Work() {
	s.work()
}

func (s Slot) Name() string {
	return s.name
}

func (s Slot) Duration() Rounds {
	return s.duration
}

func (s Slot) Cost() int {
	return s.cost
}
