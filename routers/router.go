package routers

import (
	"tools/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/test", &controllers.MainController{}, "*:Test")
	beego.Router("/json_str_dec", &controllers.MainController{}, "*:JsonStrDec")
	beego.Router("/json_file", &controllers.MainController{}, "*:JsonFile")
	beego.Router("/json_file_enhance", &controllers.MainController{}, "*:JsonFileEnhance")

	beego.Router("/go_curl", &controllers.MainController{}, "*:GoCurl")

	beego.Router("/upload", &controllers.UploadController{})
}
