package main

import (
	"fmt"
	_ "tools/routers"

	"dateparse"
	"github.com/astaxie/beego"
	"strings"
	// "runtime"
)

func main() {
	str := "2018.6.7"
	if strings.Contains(str, ".") {
		str = strings.Replace(str, ".", "/", -1)
	}
	date, err := dateparse.ParseLocal(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(date)
	}
	// runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	beego.Run()
}
