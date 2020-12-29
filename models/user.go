package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
)

type User struct {
	UserName string `form:"username"`
	Password string `form:"password"`
	PostBox  string `form:"email"`
}

//保存用户信息
func (u User) SaveUser() (int64, error) {
	u.Password = util.MD5Hash(u.Password)
	row, err := db_mysql.Db.Exec("insert into user(email, password)"+"values(?,?)", u.PostBox, u.Password)
	if err != nil {
		return -1, err
	}
	id, err := row.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

//查询用户信息
func (u User) QueryUser() (*User, error) {
	u.Password = util.MD5Hash(u.Password)
	row := db_mysql.Db.QueryRow("select email, from user where  email = ? and password = ?",
		u.PostBox, u.Password)
	err := row.Scan(&u.PostBox, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//通过邮箱查询用户信息
func QueryUserPostBox(email string) (*User, error) {
	row := db_mysql.Db.QueryRow("select email, from user where email = ?", email)
	var user User
	err := row.Scan(&user.PostBox)
	if err != nil {
		return nil, err
	}
	return &user, nil
}