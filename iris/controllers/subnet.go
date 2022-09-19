package controllers

import (
	"tools/iris/services"

	"github.com/kataras/iris/v12"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type SubnetController struct {
	BaseController
}

// @Router /subnets/clusters/cidr [get]
func (c *SubnetController) GetClusterCIDR(ctx iris.Context) {
	kubeconfig := ctx.URLParam("kubeconfig")
	ret, err := services.GetSubentService(ctx).GetClusterCIDR(kubeconfig)
	if err != nil {
		c.EchoErr(ctx, err)
		return
	}
	c.EchoJsonOk(ctx, ret)
}

// @Router /subnets/clusters/wethercontains [get]
func (c *SubnetController) WhetherContains(ctx iris.Context) {
	subnetA := ctx.URLParam("a")
	subnetB := ctx.URLParam("b")
	ret := services.GetSubentService(ctx).WhetherContains(subnetA, subnetB)
	c.EchoJsonOk(ctx, ret)
}
