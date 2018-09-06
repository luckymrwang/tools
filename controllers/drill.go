package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"unsafe"

	"github.com/astaxie/beego/httplib"
)

type DrillController struct {
	BaseController
}

func (c *DrillController) Test() {
	c.Ctx.WriteString("OK")
}

func (c *DrillController) GoCurl() {
	sql := c.GetString("sql")

	param := map[string]interface{}{
		"queryType": "SQL",
		"query":     sql,
	}

	bytesData, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)

	target := c.GetString("target", "http://localhost:8047")
	url := fmt.Sprintf("%s/query.json", target)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}

	start := time.Now()
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	end := time.Now()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println("consuming:", end.Sub(start).String())
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	c.Ctx.WriteString(*str)
}

func (c *DrillController) Jdbc() {
	sql := c.GetString("sql")

	db := map[string]string{
		"username": "test",
		"host":     "datahunter.cn",
		"password": "dh17test",
		"name":     "test",
		"fmt":      "mysql",
	}
	dbStr, err := json.Marshal(db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	start := time.Now()
	req := httplib.Post("http://jdbcdev.datahunter.cn/sql")
	req.Param("db", string(dbStr))
	req.Param("sql", sql)
	res, err := req.Response()
	end := time.Now()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("consuming:", end.Sub(start).String())
	}

	by, _ := ioutil.ReadAll(res.Body)
	c.Ctx.WriteString(string(by))
}
