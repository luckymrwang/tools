package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/test", &controllers.MainController{}, "*:Test")
	beego.Router("/json_str_dec", &controllers.MainController{}, "*:JsonStrDec")
}
