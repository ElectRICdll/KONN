package tel

import (
	"fmt"
	"konn/info"
	"reflect"
	"strings"
)

// TODO:
type EventBag struct {
	eventType reflect.Type
	eventBody Event
}

func (eb *EventBag) String() string {
	return fmt.Sprintln("<EVENTBAG>" + info.THENAME + "\n" + 
						"<VERSION>" + info.VERSION + "\n" + 
						"<TYPE>" + strings.Split(eb.eventType.String(), ".")[1] + "\n" + 
						"<BODY>" + eb.eventToSend() + "<BODYEND>")
}

func (eb *EventBag) eventToSend() string {
	return eb.String()
}

func decodeEventStr(estr string) (EventBag, error) {
	// TODO: !! should u handle this
	var res EventBag
	return res, nil
}

func decodeEvent(eb EventBag) *Event {
	event := eb.eventBody
	return &event
}

func encodeEvent(e Event) *EventBag {
	eventbag := &EventBag{reflect.TypeOf(e), e}
	return eventbag
}