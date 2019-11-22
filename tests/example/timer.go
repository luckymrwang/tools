package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case t := <-ticker.C:
			ticker.Stop()
			fmt.Println(t, "hellddddo world")
		}
	}
}
