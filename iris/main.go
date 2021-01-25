package main

import (
	"fmt"
	"tools/iris/routers"

	"github.com/kataras/iris/v12"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @host
// @BasePath /api/v1
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recover:", r)
		}
	}()
	app := iris.New()

	routers.Init(app)
	fmt.Println("continue")
	app.Run(iris.Addr(":9090"))
}

func p() string {
	panic("xxxddd")
	return "panice after"
}