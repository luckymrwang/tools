package main

import "fmt"

func main() {
	m := make(map[string]bool)
	m["A"] = true
	fmt.Println(m)
	for key := range m {
		fmt.Println(key)
	}
	delete(m, "A")
	for key, _ := range m {
		fmt.Println(key)
	}
}
