package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123"
	fmt.Println("Start...")
	switch {
	case strings.HasPrefix(str, "2"):
		fmt.Println("2")
	case strings.HasPrefix(str, "1"):
		fmt.Println("1")
	default:
		fmt.Println("default")
	}
	fmt.Println("End.")
}
