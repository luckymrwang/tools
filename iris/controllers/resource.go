package controllers

import "github.com/kataras/iris/v12"

type ResourceController struct {
	BaseController
}

func (c *ResourceController) List(ctx iris.Context) {}

func (c *ResourceController) GetResouce(ctx iris.Context) {}

func (c *ResourceController) DeleteResource(ctx iris.Context) {}
