package routers

import (
	"internerBanking/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get:HelloInternetBanking")
}
