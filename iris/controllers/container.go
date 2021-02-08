package controllers

import (
	"tools/iris/services"

	"github.com/kataras/iris/v12"
)

type ContainerController struct {
	BaseController
}

// @Summary Demo
// @Description Demoxx
// @Tags 接口
// @Accept json
// @Produce json
// @Param namespace path string true "namespace"
// @Param pod path string true "pod"
// @Param container path string true "container"
// @Success 200 {string} string	"ok"
// @Router /namespaces/{namespace}/pods/{pod}/containers/{container}/shell [get]
func (c *ContainerController) HandleExecShell(ctx iris.Context) {
	namespace := ctx.Params().Get("namespace")
	pod := ctx.Params().Get("pod")
	container := ctx.Params().Get("container")
	kubeconfig := ctx.URLParam("kubeconfig")
	sessionID, err := services.GetContainerService(ctx).ExecShell(kubeconfig, namespace, pod, container)
	if err != nil {
		return
	}
	services.GetWebSocketService(ctx).Upgrade(ctx.ResponseWriter(), ctx.Request(), sessionID)
}

// @Summary Demo
// @Description Demoxx
// @Tags 接口
// @Accept json
// @Produce json
// @Param namespace path string true "namespace"
// @Param pod path string true "pod"
// @Param container path string true "container"
// @Success 200 {string} string	"ok"
// @Router /namespaces/{namespace}/pods/{pod}/containers/{container}/copy [put]
func (c *ContainerController) CopyFromPod(ctx iris.Context) {
	namespace := ctx.Params().Get("namespace")
	pod := ctx.Params().Get("pod")
	container := ctx.Params().Get("container")
	srcPath := ctx.URLParam("src_path")
	kubeconfig := ctx.URLParam("kubeconfig")
	_, err := services.GetContainerService(ctx).CopyFromPod(kubeconfig, namespace, pod, container, srcPath)
	if err != nil {
		c.EchoErr(ctx, err)
		return
	}
	c.EchoJsonOk(ctx)
}

// @Summary Demo
// @Description Demoxx
// @Tags 接口
// @Accept json
// @Produce json
// @Param namespace path string true "namespace"
// @Param pod path string true "pod"
// @Param container path string true "container"
// @Success 200 {string} string	"ok"
// @Router /namespaces/{namespace}/pods/{pod}/containers/{container}/publish [put]
func (c *ContainerController) Publish(ctx iris.Context) {
	namespace := ctx.Params().Get("namespace")
	pod := ctx.Params().Get("pod")
	container := ctx.Params().Get("container")
	srcPath := ctx.URLParam("src_path")
	kubeconfig := ctx.URLParam("kubeconfig")
	_, err := services.GetContainerService(ctx).PublishNodeJS(kubeconfig, namespace, pod, container, srcPath)
	if err != nil {
		c.EchoErr(ctx, err)
		return
	}
	c.EchoJsonOk(ctx)
}
