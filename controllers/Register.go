package controllers

import (
	"BitcoinProject/models"
	"github.com/astaxie/beego"
)

type Register struct {
	beego.Controller
}

func (r *Register)Get()  {
	r.TplName="register.html"
}
/*
注册
*/
func (r *Register) Post() () {
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		//fmt.Println(err.Error())
		r.Ctx.WriteString("抱歉，解析数据错误")
		return
	}

	//判断该邮箱是否注册过
	_,err=user.QueryEmail()
	if err == nil{
		r.Ctx.WriteString("用户已存在")
		return
	}

	//把输入信息根据字段存入数据库的user表
	_, err = user.SaveUser()
	if err != nil {
		r.Ctx.WriteString("注册失败")
		return
	}
	r.TplName = "login.html"
}
