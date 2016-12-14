package main

import (
	"fmt"
	"strconv"
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
	//	people = make([]string, 2)
	//	people = append(people, "Boy", "Girl")
	people = []string{"Boy", "Girl"}
	Log("hello:", people, "success ", "Good")

	var a int = 10
	var b int64 = 12
	Log(b - int64(a))

	var f float64 = 0.365458
	fmt.Println(f)
	t := fmt.Sprintf("%.2f", f)
	fmt.Println(t)
	v, _ := strconv.ParseFloat(t, 64)
	fmt.Println(v)
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
