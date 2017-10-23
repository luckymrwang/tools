package main

import (
	_ "tools/routers"

	"github.com/astaxie/beego"
	// "runtime"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
}
