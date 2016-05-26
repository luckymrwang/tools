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
}
