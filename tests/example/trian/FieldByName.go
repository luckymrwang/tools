package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var m manager

	t := reflect.TypeOf(m)

	name, _ := t.FieldByName("name") // 按名称查找
	fmt.Println(name.Name, name.Type)

	age := t.FieldByIndex([]int{0, 1}) // 按多级索引查找
	fmt.Println(age.Name, age.Type)
}
