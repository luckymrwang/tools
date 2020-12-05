package controllers

import (
	"fmt"
	"tools/iris/services"

	"common.dh.cn/test"
	"github.com/kataras/iris/v12"
)

type ContainerController struct {
}

// @Summary Demo
// @Description Demoxx
// @Tags 接口
// @Accept json
// @Produce json
// @Param namespace path string true "namespace"
// @Param pod path string true "pod"
// @Param container path string true "container"
// @Param DemoParam body test.DemoParam true "container"
// @Success 200 {string} string	"ok"
// @Router /namespaces/{namespace}/pods/{pod}/containers/{container}/shell [get]
func (c *ContainerController) HandleExecShell(ctx iris.Context) {
	var demoParam test.DemoParam
	err := ctx.ReadJSON(&demoParam)
	if err != nil {
		fmt.Println("parse param error", err)
		return
	}
	namespace := ctx.Params().Get("namespace")
	pod := ctx.Params().Get("pod")
	container := ctx.Params().Get("container")
	sessionID, err := services.GetContainerService(ctx).ExecShell(namespace, pod, container)
	if err != nil {
		return
	}
	services.GetWebSocketService(ctx).Upgrade(ctx.ResponseWriter(), ctx.Request(), sessionID)
}
