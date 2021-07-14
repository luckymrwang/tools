package channel

import (
	"fmt"
	"reflect"
	"testing"
)

var rootMap = map[string]string{
	"a": "b",
	"c": "d",
}

func TestMap(t *testing.T) {
	m1 := getM()
	m2 := getM()
	m2["e"] = "f"
	if reflect.DeepEqual(m1, m2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func getM() map[string]string {
	m := make(map[string]string)
	m = rootMap
	return m
}
