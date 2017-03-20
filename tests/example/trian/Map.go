package main

import (
	"fmt"
)

func main() {
	test := make([]int, 3)
	fmt.Println(test)
	test2 := new([3]int)
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
