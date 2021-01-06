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
    //跳转到忘记密码页面
	beego.Router("/updatapwd.html",&controllers.SendCode{})
    //忘记密码发送验证码到邮箱
	beego.Router("/sendcode",&controllers.SendCode{})
    //修改密码
	beego.Router("/uppwd",&controllers.UpPwd{})
    //功能页面
    beego.Router("/home.html",&controllers.HoemController{})
    //执行方法
    beego.Router("/home",&controllers.HoemController{})
}
