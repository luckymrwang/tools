package controllers

import (
	"fmt"
	"tools/iris/services"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/kataras/iris/v12"
)

type DemoController struct {
}

type Acc struct {
	A int    `json:"a"`
	B string `json:"b"`
}

// @Router /echo [post]
func (c *DemoController) Echo(ctx iris.Context) {
	var acc Acc
	err := ctx.ReadJSON(&acc)
	if err != nil {
		fmt.Println(err)
		return
	}
	dService := services.GetDemoService(ctx)
	dService.Echo()
}
