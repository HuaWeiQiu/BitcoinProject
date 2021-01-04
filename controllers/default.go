package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
/*
 *默认显示注册页面
 */
func (c *MainController) Get() {
	c.TplName = "home.html"
}
