package main

import (
	"fmt"
	"reflect"
)

type tester interface {
	test()
	string() string
}

type data struct{}

func (*data) test()         {}
func (data) string() string { return "This is string func." }

func main() {
	var d data

	// var t tester = d

	// 错误: data does not implement tester
	// (test method has pointer receiver)

	var t tester = &d

	t.test()
	fmt.Println(t.string())

	var a interface{}
	a = b()
	if reflect.ValueOf(a).IsNil() {
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
}

func b() *int {
	return nil
}
