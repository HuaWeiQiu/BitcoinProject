package models

import (
	"BitcoinProject/db_mysql"
	"BitcoinProject/utils"
)

type User struct {
	Id        int    `form:"id"`
	Email     string `form:"email"`
	Password  string `form:"password"`
	TimeStamp int64  `form:"timestamp"`
	Code      string `form:"code"`
}

/*
保存用户信息的方法：保存用户信息到数据库中
*/
func (u User) SaveUser() (int64, error) {
	//1.密码脱敏处理
	u.Password = utils.Md5Hash(u.Password)
	//2.执行数据库操作
	row, err := db_mysql.Db.Exec("insert into user (email,password) values(?,?)", u.Email, u.Password)
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
	row := db_mysql.Db.QueryRow("select email from user where email = ? and password = ?", u.Email, u.Password)
	err := row.Scan(&u.Email)
	if err != nil {
		return nil, err
	}
	return &u, err
}

//通过邮箱查询用户信息
func (u User) QueryByEmail(email string) (*User, error) {
	//u.Password = utils.Md5Hash(u.Password)
	row := db_mysql.Db.QueryRow("select email from user where email = ?", email)
	var user User
	err := row.Scan(&user.Email)
	if err != nil {
		return nil, err
	}
	return &user, err
}
/*
更新(插入)数据
*/
func (u User)Updata(email string)(int64,error)  {
	rs, err := db_mysql.Db.Exec("update user set timestamp=? ,code=? where email=?",u.TimeStamp,u.Code,email)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}
/*
修改密码
*/
func (u User)UpPwd(email string)(int64,error)  {
	u.Password = utils.Md5Hash(u.Password)
	rs, err := db_mysql.Db.Exec("update user set password=? where email=?",u.Password,email)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}
/*
如果验证码超过时长就删除验证码
*/
func (u User)DelCode(email string)(int64,error)  {
	rs, err := db_mysql.Db.Exec("update user set code=null where email=?",u.Code,email)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}
//根据查询条件(email)进行数据查询
func (u User) QueryByEmails(email string) (*User, error) {
	//u.Password = utils.Md5Hash(u.Password)
	row := db_mysql.Db.QueryRow("select email, timestamp ,code from user where email=?", email)
	var user User
	err := row.Scan(&user.Email,&user.TimeStamp,&user.Code)
	if err != nil {
		return nil, err
	}
	return &user, err
}