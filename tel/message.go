package tel

// TODO: maybe?
type data interface {
	pushData(data DataBag) error
	pullData(data DataBag) error
}

type message interface {
	sendMeesage(message string, des string) error
	recvMessage(message string) error
}