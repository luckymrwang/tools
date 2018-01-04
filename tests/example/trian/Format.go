package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"年龄"`
}

func main() {
	s := []Student{Student{Name: "Tom", Age: 18}, Student{Name: "Cat", Age: 20}}

	fmt.Println(ModelToArrayP(s))
}

type P map[string]interface{}

func ModelToArrayP(o interface{}) []P {
	var array = []P{}
	oo := reflect.ValueOf(o)
	if oo.Kind() != reflect.Slice {
		return array
	} else {
		for i := 0; i < oo.Len(); i++ {
			array = append(array, ModelToP(oo.Index(i).Interface()))
		}
	}

	return array
}

func ModelToP(o interface{}) P {
	info := P{}
	if o == nil {
		return info
	} else {
		s := reflect.ValueOf(o)
		if s.Kind() == reflect.Ptr {
			s = s.Elem()
		}

		for i := 0; i < s.NumField(); i++ {
			f := s.Type().Field(i)
			key := f.Tag.Get("json")
			if key == "" || key == "-" {
				continue
			}
			value := s.Field(i).Interface()
			info[key] = value
		}
		return info
	}
}
