package tel

import (
	. "konn/constants"
	. "konn/tel/events"
)

// TODO:
type EventBag struct {
	eventType string
	eventBody string
}

func (eb *EventBag) String() string {
	return EVENTBAG + MODULE_NAME + "\n" +
		THE_VERSION + VERSION + "\n" +
		EVENTTYPE + eb.eventType + "\n" +
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
	return &EventBag{e.Name, e.Message}
}
