package main

import (
	"fmt"
)

type data struct{}

func (this *data) Error() string { return "" }

func bad() bool {
	return true
}

func test() error {
	var p *data = nil
	if bad() {
		return p
	}
	return nil
}

func main() {
	var e error = test()
	if e == nil {
		fmt.Println("e is nil")
	} else {
		fmt.Println("e is not nil")
	}
}
