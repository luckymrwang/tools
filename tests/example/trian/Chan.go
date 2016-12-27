package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 4)
	for i := 0; i < 4; i++ {
		go func(chan int, int) {
			time.Sleep(time.Duration(i) * time.Second)
			ch <- 1
		}(ch, i)
	}

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("fjfj")
}
