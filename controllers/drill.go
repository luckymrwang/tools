package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
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
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	c.Ctx.WriteString(*str)
}
