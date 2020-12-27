package models

import (
	"BitcoinProject/db_mysql"
	"BitcoinProject/utils"
)

type User struct {
	Id       int    `form:"id"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
/*
保存用户信息的方法：保存用户信息到数据库中
*/
func (u User) SaveUser() (int64, error) {
	//1.密码脱敏处理
	u.Password = utils.Md5Hash(u.Password)
	//2.执行数据库操作
	row, err := db_mysql.Db.Exec("insert into user (email,password) values(?,?)",u.Email,u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, nil
	}
	return id, err
}

/*
查询数据
*/
func (u User) QueryUser() (*User, error) {
	//1、密码脱敏
	u.Password = utils.Md5Hash(u.Password)
	row := db_mysql.Db.QueryRow("select username from user where username = ? and password =?", u.Email, u.Password)
	err := row.Scan(&u.Email)
	if err != nil {
		return nil, err
	}
	return &u, err
}