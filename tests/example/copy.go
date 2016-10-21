package main

import (
	"fmt"
)

func main() {
	var slice1 map[string]string

	slice1 = make(map[string]string)
	slice1["hello"] = "boy"
	fmt.Println("slice1:", slice1)

	var a = [5]int{1, 2, 3, 4, 5}
	slice := a[1:3]
	fmt.Println(cap(slice))
}
