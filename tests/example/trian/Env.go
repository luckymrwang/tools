package main

import (
	"fmt"
	"os"
)

func main() {
	//	os.Setenv("mode", "EE")
	fmt.Println("mode", os.Getenv("mode"))

	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2
	delete(m, "h")
	fmt.Println(m)

}
