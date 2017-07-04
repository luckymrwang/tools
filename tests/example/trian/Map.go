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
	test2 := new(map[string]bool)
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

	// map
	fmt.Println(data)
	fmt.Println(Add(data))
	fmt.Println(data)

	// slice
	fmt.Println(keys)
	fmt.Println(ChangeSlice(keys))
	fmt.Println(keys)

	// array
	arr := [3]string{"a", "b", "c"}
	fmt.Println(arr)
	fmt.Println(ChangeArr(arr))
	fmt.Println(arr)
}

func Add(m map[string]float64) map[string]float64 {
	for key, _ := range m {
		m[key] += 1
	}

	return m
}

func ChangeSlice(s []string) []string {
	s[0] = "abc"

	return s
}

func ChangeArr(s [3]string) [3]string {
	s[0] = "abc"

	return s
}
