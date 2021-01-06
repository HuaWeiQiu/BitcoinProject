package main

import (
	"BitcoinProject/db_mysql"
	_ "BitcoinProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.ConnectDB()
	//设置静态资源路径
	//hash:=utils.GetBestBlockHash()
	//fmt.Println(hash)
	//heighthash,_:=utils.GetBlockHash(0)
	//fmt.Println(heighthash)
	//balances,_:=utils.GetBalances()
	//fmt.Println(balances)
	//chaintips,_:=utils.GetChaintips()
	//fmt.Println(chaintips[0])
	//difficult:=utils.GetDifficult()
	//fmt.Println(difficult)
	//memoryinfo,_:=utils.GetMemoryInfo()
	//fmt.Println(memoryinfo)
	//blockchianinfo,_:=utils.GetBlockChainInfo()
	//fmt.Println(blockchianinfo)
	//blockcount:=utils.GetBlockCount()
	//fmt.Println(blockcount)
	//rpcinfo,_:=utils.GetRpcInfo()
	//fmt.Println(rpcinfo)
	//logging,_:=utils.Logging()
	//fmt.Println(logging)
	//uptime:=utils.Uptime()
	//fmt.Println(uptime)
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run("127.0.0.1:8080")
}
