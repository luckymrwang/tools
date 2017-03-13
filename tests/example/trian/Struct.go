package main

import "fmt"

type Human struct {
	name  string
	age   int
	favor []string
	grade map[string]int
}

func main() {
	jane := Human{
		name: "jane",
		age:  10,
		favor: []string{
			"football",
			"basketball"},
		grade: map[string]int{
			"语文": 98, "数学": 100}}
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println(jane)
}
