package main

import (
	"fmt"
	"time"
)

var m map[string]string

func main() {
	m = make(map[string]string)
	m["a"] = "abc"

	for i := 0; i < 10; i++ {
		go func(j int) {
			time.Sleep(time.Duration(j) * time.Second)
			fmt.Println(getRet())
		}(i)
	}
	time.Sleep(time.Second * 5)
}

func getRet() string {
	return m["a"]
}
