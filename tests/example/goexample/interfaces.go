package main

import (
	"fmt"
)

func main() {
	test("test", 1)
	c := "abcd"
	fmt.Println(c[0])
}

func test(a string, b ...interface{}) {
	fmt.Println(a, b[0])
}
