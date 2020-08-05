package main

import (
	"fmt"
	"reflect"
)

type R struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func main() {
	r := R{
		A: "cccc",
		B: 222,
	}
	mutable := reflect.ValueOf(&r).Elem()
	fmt.Println(mutable.FieldByName("A").Interface())

	fmt.Println(r)
}
