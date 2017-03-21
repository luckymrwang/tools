package main

import (
	"fmt"
)

func main() {
	test := make(map[string]bool)
	if test == nil {
		fmt.Println("test is nil")
	}
	fmt.Println(test["index"])
	fmt.Println(test)
	test2 := new([]int)
	if *test2 == nil {
		fmt.Println("test2 is nil")
	}
	fmt.Println(test2)

	data := map[string]float64{}
	keys := []string{"a", "b", "c"}

	var val float64 = 1

	for _, key := range keys {
		v, ok := data[key]
		fmt.Println(v, ok)
		fmt.Println(data[key])
		data[key] += val
	}

	fmt.Println(data)
}
