package routers

import (
	"go-weibo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
