package tel

import (
	. "konn/constants"
	. "konn/tel/events"
	"reflect"
	"strings"
)

// TODO:
type EventBag struct {
	eventType reflect.Type
	eventBody string
}

func (eb *EventBag) String() string {
	return EVENTBAG + MODULE_NAME + br +
		THE_VERSION + VERSION + br +
		EVENTTYPE + strings.Split(eb.eventType, ".")[1] + br +
		BODY +
		eb.eventToSend() +
		BODYEND
}

func (eb *EventBag) eventToSend() string {
	return eb.eventBody
}

func decodeMessage(estr string) (EventBag, error) {
	// TODO: !! should u handle this
	var res EventBag
	return res, nil
}

// func DecodeEvent(eb EventBag) *Event {
// 	e := &basicEvent{}
// 	return e
// }

func EncodeEvent(e *Event) *EventBag {
	return &EventBag{e.Type, e.Message}
}
