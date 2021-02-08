package routers

import (
	"tools/iris/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func DeploymentHub(party iris.Party) {
	party.Put("/namespaces/{namespace}/deployments/{deployment}/inject", hero.Handler(new(controllers.DeploymentController).Inject))
}
