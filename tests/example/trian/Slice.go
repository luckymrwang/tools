package main

import (
	"fmt"
)

func main() {
	s := make([]string, 0)
	fmt.Println(s)
	for _, v := range s {
		fmt.Println("hello", v)
	}

}
