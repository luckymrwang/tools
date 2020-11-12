package main

import (
	"fmt"
	"tools/iris/routers"

	"github.com/kataras/iris/v12"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recover:", r)
			return
		}
	}()
	app := iris.New()

	routers.Init(app)
	fmt.Println("continue")
	app.Run(iris.Addr(":9090"))
}

func p() {
	panic("xxxddd")
}
