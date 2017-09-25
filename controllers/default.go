package controllers

import (
	"github.com/astaxie/beego"
)

// MainController ...
type MainController struct {
	beego.Controller
}

// Get ...
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// HelloInternetBanking ...
func (c *MainController) HelloInternetBanking() {
	c.Data["WebSite"] = "My web site"
	c.Data["Email"] = "My email"
	c.Data["EmailName"] = "My email name"
	c.Data["Id"] = c.Ctx.Input.Param(":id")

	c.TplName = "default/hello-internetBanking.tpl"
}
