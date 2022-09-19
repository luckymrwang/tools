package routers

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/kataras/iris/v12/middleware/pprof"
	irisrecover "github.com/kataras/iris/v12/middleware/recover"

	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/swaggo/swag"
)

func Init(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,          // allows browser send cookie to service
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Acceppt", "Content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	app.Use(irisrecover.New())
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
	config := &swagger.Config{
		URL: "http://localhost:8080/swagger/doc.json", //The url pointing to API definition
	}
	p := pprof.New()
	app.Any("/debug/pprof", p)
	app.Any("/debug/pprof/{action:path}", p)
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	app.Get("/v2/api-docs", func(ctx iris.Context) {
		doc, err := swag.ReadDoc()
		if err != nil {
			panic(err)
		}
		_, _ = ctx.Write([]byte(doc))
	})
	app.Get("/regionOne/iapps-service.yaml", iappserverYamlHandler)
	hubBus := app.Party("/v1", crs).AllowMethods(iris.MethodOptions)
	Hub(hubBus)

	//app.Use(iris.FromStd(func(w http.ResponseWriter, r *http.Request) {
	//	println("Request path: " + r.URL.Path)
	//}))
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
	DeploymentHub(party)
	WebsocketHub(party)
	SubnetHub(party)
}
