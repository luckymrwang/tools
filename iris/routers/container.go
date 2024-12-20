package routers

import (
	"tools/iris/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
)

func ContainerHub(party iris.Party) {
	party.Get("/namespaces/{namespace}/pods/{pod}/containers/{container}/shell", hero.Handler(new(controllers.ContainerController).HandleExecShell))
	party.Put("/namespaces/{namespace}/pods/{pod}/containers/{container}/copy", hero.Handler(new(controllers.ContainerController).CopyFromPod))
	party.Put("/namespaces/{namespace}/pods/{pod}/containers/{container}/publish", hero.Handler(new(controllers.ContainerController).Publish))
}
