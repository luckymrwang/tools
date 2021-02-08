package controllers

import (
	"tools/iris/services"

	"github.com/kataras/iris/v12"
)

type DeploymentController struct {
	BaseController
}

// @Summary Deploy
// @Description Deploy
// @Tags 接口
// @Accept json
// @Produce json
// @Param namespace path string true "namespace"
// @Param deployment path string true "deployment"
// @Param container path string true "container"
// @Success 200 {string} string	"ok"
// @Router /namespaces/{namespace}/deployments/{deployment}/inject [put]
func (c *DeploymentController) Inject(ctx iris.Context) {
	namespace := ctx.Params().Get("namespace")
	deployment := ctx.Params().Get("deployment")
	srcPath := ctx.URLParam("src_path")
	kubeconfig := ctx.URLParam("kubeconfig")
	_, err := services.GetDeploymentService(ctx).InjectSidecar(kubeconfig, namespace, deployment, srcPath)
	if err != nil {
		c.EchoErr(ctx, err)
		return
	}
	c.EchoJsonOk(ctx)
}
