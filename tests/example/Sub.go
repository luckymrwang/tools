package main

import (
	"fmt"
	"time"
)

func main() {
	d := 0
	start := time.Now()
	for i := 0; i < 800000000; i++ {
		d += i
	}

	fmt.Println("d:", d)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
