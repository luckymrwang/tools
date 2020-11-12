package controllers

import (
	"fmt"
	"tools/iris/services"

	"github.com/kataras/iris/v12"
)

type ContainerController struct {
}

func (c *ContainerController) HandleExecShell(ctx iris.Context) {
	namespace := ctx.Params().Get("namespace")
	pod := ctx.Params().Get("pod")
	container := ctx.Params().Get("container")
	fmt.Println(ctx.GetHeader("test"))
	sessionID, err := services.GetContainerService(ctx).ExecShell(namespace, pod, container)
	if err != nil {
		return
	}
	services.GetWebSocketService(ctx).Upgrade(ctx.ResponseWriter(), ctx.Request(), sessionID)
}
