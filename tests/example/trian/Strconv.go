package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "100a"
	b, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}
