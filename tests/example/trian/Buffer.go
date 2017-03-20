package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	var buffer bytes.Buffer

	t := time.Now()
	for i := 0; i < 10000000; i++ {
		buffer.WriteString("hello")
	}

	fmt.Println(time.Since(t))

	t1 := time.Now()
	s := ""
	var buf []byte
	buf = append(buf, []byte(s)...)
	for j := 0; j < 10000000; j++ {
		buf = append(buf, []byte("hello")...)
	}

	fmt.Println(time.Since(t1))
}
