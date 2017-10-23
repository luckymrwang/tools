package controllers

import (
	"fmt"
	//	"io/ioutil"
	"time"

	//	"github.com/astaxie/beego/httplib"

	"tools/utils"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}

func (c *MainController) Test() {
	c.Ctx.WriteString("OK")
}

func (c *MainController) JsonStrDec() {
	utils.JsonStream()
}

func (c *MainController) JsonFile() {
	utils.JsonStreamFile()
}

func (c *MainController) JsonFileEnhance() {
	_, ms := utils.JsonStreamFileEnhance()
	for _, val := range ms {
		str := fmt.Sprintf("%s : %d \n", val.Key, val.Value)
		c.Ctx.WriteString(str)
	}
}

func (c *MainController) GoCurl() {
	//	cnt := 10000
	//	for i := 1; i <= 620000; i += cnt {
	//		str := fmt.Sprintf("http://119/bi-agent2/api/maintain/fix_by_device_id?appid=2005&start=%d&end=%d", i, i+cnt)

	//		fmt.Println(str)
	//		req, err := httplib.Get(str).SetTimeout(10*time.Minute, 10*time.Minute).Response()
	//		if err != nil {
	//			fmt.Println(err)
	//		}

	//		by, _ := ioutil.ReadAll(req.Body)
	//		c.Ctx.WriteString(string(by))
	//	}

	start := "2017-02-01"
	end := "2017-04-01"
	t1, _ := time.Parse("2006-01-02", start)
	t2, _ := time.Parse("2006-01-02", end)

	for t := t1; t2.Sub(t) >= 0; t = t.AddDate(0, 1, 0) {
		//		fmt.Println(t.Format("2006-01-02"), time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02"))
		sd := t.Format("2006-01-02")
		ed := time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
		str := fmt.Sprintf("http://52/bi-agent/task/crontab/fix_key_deadline?start_date=%s&end_date=%s", sd, ed)

		fmt.Println(str)
		//		req, err := httplib.Get(str).SetTimeout(10*time.Minute, 10*time.Minute).Response()
		//		if err != nil {
		//			fmt.Println(err)
		//		}

		//		by, _ := ioutil.ReadAll(req.Body)
		//		c.Ctx.WriteString(string(by))
	}
}
