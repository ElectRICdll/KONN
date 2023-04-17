package tel

import (
	"konn/constants"
	"reflect"
	"strings"
)

// TODO:
type EventBag struct {
	eventType reflect.Type
	eventBody string
}

func (eb *EventBag) String() string {
	return "<EVENTBAG>" + constants.THENAME + "\n" + 
		"<VERSION>" + constants.VERSION + "\n" + 
		"<TYPE>" + strings.Split(eb.eventType.String(), ".")[1] + "\n" + 
		"<BODY>" + eb.eventToSend() + "<BODYEND>"
}

func (eb *EventBag) eventToSend() string {
	return eb.eventBody
}

func decodeEventStr(estr string) (EventBag, error) {
	// TODO: !! should u handle this
	var res EventBag
	return res, nil
}

// func DecodeEvent(eb EventBag) *Event {
// 	e := &basicEvent{}
// 	return e
// }

func EncodeEvent(e Event) *EventBag {
	eventbag := &EventBag{reflect.TypeOf(e), e.String()}
	return eventbag
}