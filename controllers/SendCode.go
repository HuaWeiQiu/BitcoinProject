package controllers

import (
	"BitcoinProject/Email"
	_ "BitcoinProject/Email"
	"BitcoinProject/models"
	"BitcoinProject/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type SendCode struct {
	beego.Controller
}

func (s *SendCode)Get()  {
	s.TplName="uppwd.html"
}

//邮件主题为"Hello"
var subject = "Hello by golang gomail from 1987752122@qq.com"

//验证码
var code = utils.GenCode(4)

// 邮件正文
var body = `<h3 style ="color:red">尊敬的用户您好</h3>` +
	`<span style = "color:blue">您的验证码是</span>` + code+",验证码有效期为30分钟"

func (s *SendCode) Post() {
	fmt.Println("是否为ajax请求",s.IsAjax())
	var user models.User
	var codes models.SendCode
	//err := s.ParseForm(&user)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(user.Email)

	data:=s.Ctx.Input.RequestBody
	json.Unmarshal(data,&codes)
	fmt.Println(codes)
	fmt.Println(code)
	if codes.Codeinfo!="获取验证码"{
		return
	}
	//查询该用户是否注册过
	_, err := user.QueryByEmail(codes.Email)
	if err != nil {
		s.Data["json"]="该邮箱尚未注册,请先注册"
		s.ServeJSON()
		fmt.Println(err.Error())
		return
	}
	//定义收件人
	mailTo := []string{
		codes.Email,
		//"aaa2398546675@163.com",
		//"qhw1987752122@163.com",
		//"930987934@qq.com",
	}
	//发送邮件
	err = Email.SendMail(mailTo, subject, body)
	if err != nil {
		fmt.Println("send fail")
		return
	}
	user=models.User{
		TimeStamp:time.Now().Unix(),
		Code:code,
	}
	_,err=user.Updata(codes.Email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("send successfully")
	s.Data["json"]="发送成功"
	s.ServeJSON()
}

