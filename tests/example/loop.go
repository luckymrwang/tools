package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	c := time.Tick(10 * time.Second)
	for {
		select {
		case <-c:
			fmt.Println(i + 1)
		}
	}
}
