package controllers

import (
	"fmt"
	"log"
)

type UploadController struct {
	BaseController
}

func (c *UploadController) Get() {
	c.TplName = "upload.tpl"
	c.Render()
}

func (c *UploadController) Post() {
	f, h, err := c.GetFile("uploadname")
	fmt.Println("auth", c.GetString("auth"))
	fmt.Println(f, h)
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	c.Ctx.WriteString("ok")
}
