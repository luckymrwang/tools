package main

import (
	"fmt"
)

func main() {
	//	fmt.Printf("Hello, world\n")
	str := "Hello"
	n := len(str)
	for i := 0; i < n; i++ {
		//		ch := str[i] // 依据下标取字符串中的字符,类型为byte
		ch := str[i:i]

		fmt.Println(i, ch)

	}
}
