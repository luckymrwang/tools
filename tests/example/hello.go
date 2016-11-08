package main

import (
	"fmt"
)

func main() {
	//	fmt.Printf("Hello, world\n")
	str := "Hello"
	c := []byte(str)
	fmt.Println(c)
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符,类型为byte
		fmt.Println(i, ch)

	}

	var people []string
	//	people := make([]string, 2)
	//	people = append(people, "Boy", "Girl")
	//	people := []string{"Boy", "Girl"}
	Log("hello:", people)
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
