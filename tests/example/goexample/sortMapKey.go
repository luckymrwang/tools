package main

import (
	"fmt"
	"sort"
)

func main() {
	// To create a map as input
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
