package main

import (
	"github.com/astaxie/beego"
	_ "hello/routers"
	// "runtime"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
}
