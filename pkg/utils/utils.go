package utils

import "reflect"

func IsNilOrEmpty(i interface{}) bool {
	switch v := i.(type) {
	case map[string]interface{}:
		return len(v) == 0
	default:
		return i == nil || reflect.ValueOf(i).IsNil()
	}
}
