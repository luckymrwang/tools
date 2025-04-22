package controllers

import (
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"tools/iris/services"

	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"github.com/kataras/iris/v12"
)

type DemoController struct {
	BaseController
}

type Acc struct {
	A int    `json:"a"`
	B string `json:"b"`
}

type U struct {
	Object map[string]interface{}
}

type NetApi struct {
	Net *U `json:"net"`
	unstructured.Unstructured
}

func (u *U) MarshalJSON() ([]byte, error) {
	// 忽略Object字段
	type alias U
	return json.Marshal((*alias)(u))
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

// @Router /get [get]
func (c *DemoController) Get(ctx iris.Context) {
	var net *NetApi
	net = &NetApi{Net: &U{Object: map[string]interface{}{"a": "b", "c": "d"}}}
	c.EchoJsonOk(ctx, net)
}

// @Router /pu [post]
func (c *DemoController) Pu(ctx iris.Context) {
	var net *NetApi
	err := ctx.ReadJSON(&net)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.EchoJson(ctx, net)
}
