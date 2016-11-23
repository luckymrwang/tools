package controllers

import (
	"hello/tools"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test() {
	c.Ctx.WriteString("OK")
}

func (c *MainController) JsonStrDec() {
	tools.JsonStream()
}

func (c *MainController) JsonFile() {
	tools.JsonStreamFile()
}
