package routers

import (
	"tools/iris/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func DemoHub(party iris.Party) {
	party.Get("/namespaces/helm", hero.Handler(new(controllers.DemoController).Helm))
}
