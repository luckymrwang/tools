package main

import (
	"fmt"
	"time"
	//	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
	fmt.Println("start sleeping...")
	time.Sleep(time.Second * 1)
	fmt.Println("end sleep.")
}
