package main

import "fmt"

type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type A struct {
}

func (a A) Read() {
}

var _ Reader = &A{} // 编译通过，确保A实现了 Reader 接口
var _ Writer = &A{} // 编译通不过，A没有实现 Writer 接口

func main() {
	fmt.Println("aaa")
}
