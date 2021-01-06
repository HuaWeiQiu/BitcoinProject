package controllers

import (
	"BitcoinProject/models"
	"github.com/astaxie/beego"
)

type Login struct {
	beego.Controller
}

func (l *Login) Get() {
	l.TplName = "login.html"
}

func (l *Login) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户信息解析失败")
		return
	}
	//判断该邮箱是否存在
	_,err=user.QueryByEmail(user.Email)
	if err != nil{
		l.Ctx.WriteString("用户尚未注册")
		return
	}
	//查询数据库用户信息
	u, err := user.QueryUser()
	if err != nil {
		//fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，登陆失败请重试")
		return
	}
	//登陆成功，跳转项目核心功能页面
	l.Data["Email"] = u.Email
	l.TplName = "home.html"
}
