package controllers

import (
	"tools/iris/services"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/kataras/iris/v12"
)

type DemoController struct {
}

func (c *DemoController) Echo(ctx iris.Context) {
	dService := services.GetDemoService(ctx)
	dService.Echo()
}
