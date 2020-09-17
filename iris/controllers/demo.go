package controllers

import (
	"tools/iris/services"

	"github.com/kataras/iris/v12"
)

type DemoController struct {
}

func (c *DemoController) Echo(ctx iris.Context) {
	dService := services.GetDemoService(ctx)
	dService.Echo()
}
