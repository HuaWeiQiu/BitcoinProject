package routers

import (
	"BitcoinProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    //注册
    beego.Router("/register",&controllers.Register{})
    //跳转到注册页面
    beego.Router("/register.html",&controllers.Register{})
    //注册后跳转登陆界面
    beego.Router("/login",&controllers.Login{})
    //在注册页面已有账号直接跳转到登陆界面
    beego.Router("/login.html",&controllers.Login{})
}
