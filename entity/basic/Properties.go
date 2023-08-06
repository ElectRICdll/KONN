package basic

import (
	"fmt"
	"reflect"
)

/*
The Properties Class provides a standard of entities' basical properties.
*/
type Properties struct {
	Health    int
	Armor     int
	Scout     int
	AntiScout int
	Flex      int
}

func (p Properties) String() string {
	var result string
	typeof := reflect.TypeOf(p)
	valueof := reflect.ValueOf(p)
	for i := 0; i < typeof.NumField(); i++ {
		result += fmt.Sprintf("%s: %v", typeof.Field(i), valueof.Field(i).Interface())
	}

	return result
} // TODO: ?
