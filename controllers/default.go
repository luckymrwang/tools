package controllers

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego/httplib"

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

func (c *MainController) GoCurl() {
	cnt := 10000
	for i := 1; i <= 620000; i += cnt {
		str := fmt.Sprintf("http://119.29.217.39/bi-agent2/api/maintain/fix_by_device_id?appid=2005001001&start=%d&end=%d", i, i+cnt)

		fmt.Println(str)
		req, err := httplib.Get(str).SetTimeout(10*time.Minute, 10*time.Minute).Response()
		if err != nil {
			fmt.Println(err)
		}

		by, _ := ioutil.ReadAll(req.Body)
		c.Ctx.WriteString(string(by))
	}
}
