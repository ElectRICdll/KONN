package gameplay

var subManager *SubManager

type SubManager struct {
	idPool [10000]int
	subNum int
}

func InstanceOfSubManager() *SubManager {
	if subManager == nil {
		subManager = &SubManager{
			idPool: [10000]int{},
			subNum: 0,
		}
	}
	return subManager
}

func (s *SubManager) IDGenerate() {
	//snowflake.Epoch
}
