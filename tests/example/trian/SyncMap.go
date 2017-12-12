package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	m.Store("one", 1)
	m.Store("one", 2)

	if v, ok := m.Load("one"); ok {
		fmt.Println(v)
	}

	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)

		return true
	})
}
