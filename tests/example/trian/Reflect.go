package main

import (
	"fmt"
	"reflect"
)

func main() {
	typeShow()
	pointer()
	mapReflect()
}

func pointer() {
	x := 100

	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println(tx, tp, tx == tp)

	fmt.Println(tx.Kind(), tp.Kind())
	fmt.Println(tx == tp.Elem())
}

func mapReflect() {
	abc := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println(reflect.ValueOf(abc))
	keys := reflect.ValueOf(abc).MapKeys()

	fmt.Println(keys) // [a b c]
}

type X int
type Y int

func typeShow() {
	var A X = 100
	t := reflect.TypeOf(A)

	fmt.Println(t)
	fmt.Println(t.Name(), t.Kind())

	var a, b X = 100, 200
	var c Y = 300

	ta, tb, tc := reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c)

	fmt.Println(ta == tb, ta == tc)
	fmt.Println(ta.Kind() == tc.Kind())
}

func recomp() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	m := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))

	fmt.Println(a, m)
}
