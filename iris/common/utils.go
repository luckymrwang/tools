package common

import "reflect"

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	vi := reflect.ValueOf(v)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
