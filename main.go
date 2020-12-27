package main

import (
	"BitcoinProject/db_mysql"
	_ "BitcoinProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.ConnectDB()
	beego.Run()
}

