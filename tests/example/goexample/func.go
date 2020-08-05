package main

import "fmt"

func main() {
	a, b := hello()
	fmt.Println(a, b)
}

func hello() (a string, b int) {
	a = "hello"
	d := 2
	c := "boy"

	return c, d
}
