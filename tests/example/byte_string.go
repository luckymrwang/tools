package main

import (
	"fmt"
)

func main() {
	s1 := "abcd"
	b1 := []byte(s1)
	fmt.Println(b1) // [97 98 99 100]

	s2 := "中文"
	b2 := []byte(s2)
	fmt.Println(b2) // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成

	r := []rune(s2)
	fmt.Println(r) // [20013 25991], 每个字一个数值

	str := "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch) // ch的类型byte
	}

	for i, ch := range str {
		fmt.Println(i, ch) // ch的类型为rune
	}
}
