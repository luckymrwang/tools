package main

import (
	"fmt"
	"time"
)

func main() {
	start := "2017-02-17"
	end := "2017-03-13"
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)

	for t := t1; t2.Sub(t) >= 0; t = t.Add(24 * time.Hour) {
		fmt.Println(t.Format("2006-01-02"))
	}
}
