package routers

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"time"
)

func Init(app *iris.Application) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
			return
		}
	}()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,          // allows browser send cookie to service
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Acceppt", "Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
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
	hubBus := app.Party("/v1", crs).AllowMethods(iris.MethodOptions)
	Hub(hubBus)

	app.Use(iris.FromStd(func(w http.ResponseWriter, r *http.Request) {
		println("Request path: " + r.URL.Path)
	}))
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

func Hub(party iris.Party) {
	DemoHub(party)
	ContainerHub(party)
	WebsocketHub(party)
}

func p() {
	panic("xxxddd")
}
