package main

import (
	_ "hello/routers"
	"github.com/astaxie/beego"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
}

