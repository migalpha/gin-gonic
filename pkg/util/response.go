package util

import (
	"reflect"
)

// ParseArray for return empty arrays in responses
func ParseArray(a interface{}) (res interface{}) {
	if a == nil || reflect.ValueOf(a).IsNil() {
		res = []string{}
	} else {
		res = a
	}
	return
}
