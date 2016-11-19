package main

import (
	"fmt"
)

func main() {
	var slice1 map[string]string

	slice1 = make(map[string]string)
	slice1["hello"] = "boy"
	slice1["hello"] = "boy"
	slice1["he"] = "girl"
	slice1["he"] = "girl"
	fmt.Println("slice1:", slice1)

	var a = [5]int{1, 2, 3, 4, 5}
	slice := a[1:3]
	fmt.Println(cap(slice))

	x := 7
	y := 8

	max(x, y)
	fmt.Println(y)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
