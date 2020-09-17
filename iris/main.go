package main

import (
	"tools/iris/routers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	routers.Init(app)
	app.Run(iris.Addr(":9090"))
}
