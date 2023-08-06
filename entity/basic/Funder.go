package basic

type Funder struct {
	PerSecond int
	IsWorking bool
}

func (f *Funder) Start() {
	f.IsWorking = true
}

func (f *Funder) Halt() {
	f.IsWorking = false
}
