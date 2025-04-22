package routers

import (
	"tools/iris/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func DemoHub(party iris.Party) {
	party.Post("/echo", hero.Handler(new(controllers.DemoController).Echo))
	party.Get("/get", hero.Handler(new(controllers.DemoController).Get))
	party.Post("/pu", hero.Handler(new(controllers.DemoController).Pu))
}
