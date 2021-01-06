package controllers

import (
	"BitcoinProject/models"

	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UpPwd struct {
	beego.Controller
}

func (u *UpPwd) Post() {
	fmt.Println("是否为ajax请求", u.IsAjax())
	var user models.User
	var codes models.SendCode
	data := u.Ctx.Input.RequestBody
	json.Unmarshal(data, &codes)
	fmt.Println(codes)
	fmt.Println(code)
	if codes.Code == code {
		now := time.Now().Unix()
		fmt.Println("xxxxxxx")
		users, _ := user.QueryByEmails(codes.Email)
		fmt.Println(users)
		if now-users.TimeStamp < 1000*60*3 {
			user = models.User{
				Password: codes.Password,
			}
			_, err := user.UpPwd(codes.Email)
			if err != nil {
				fmt.Println(err.Error())
				u.Ctx.WriteString("修改密码错误")
				return
			}
			u.Data["json"] = "修改密码成功"
			u.ServeJSON()
		}
	} else {
		_,err:=user.DelCode(codes.Email)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
}
