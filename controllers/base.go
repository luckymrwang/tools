package controllers

import (
	"github.com/astaxie/beego"
)

func init() {}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) JsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) RespMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["errcode"] = msgno
	out["errmsg"] = msg

	this.JsonResult(out)
}
