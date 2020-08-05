package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	// app.Get("/", before, mainHandler, after)
	app.Use(func(ctx iris.Context) {
		time.Now()
		ctx.Next()
	})
	t := time.Now()
	app.Done(func(ctx iris.Context) {
		fmt.Println("finished：", time.Since(t))
		ctx.Next()
	})
	app.Get("/regionOne/iapps-service.yaml", iappserverYamlHandler)

	app.Run(iris.Addr(":9090"))
}

func iappserverYamlHandler(ctx iris.Context) {
	path := "./iapps-server.yaml"
	c, err := ioutil.ReadFile(path)
	if err != nil {
		println(err)
	}
	ctx.WriteString(string(c))
}

func loginNameHandler(ctx iris.Context) {
	name := ctx.Params().Get("name")
	println(name)
	ctx.Next()
}

func loginHandler(ctx iris.Context) {
	println("login")
	ctx.Next()
}

func before(ctx iris.Context) {
	println("before")
	ctx.Next() //继续执行下一个handler，这本例中是mainHandler
}

func mainHandler(ctx iris.Context) {
	println("mainHandler")
	ctx.Next()
}

func after(ctx iris.Context) {
	println("after")
	ctx.Next()
}
