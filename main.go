package main

import (
	_ "BitcoinProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//设置静态资源路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}
