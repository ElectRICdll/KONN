package events

import "konn/entity/prop"

type DeactivateEvent struct {
	who prop.Substantive
}

// TODO
func NewDeactivateEvent() {

}

func (e DeactivateEvent) ToMessage() string {
	return ""
}
