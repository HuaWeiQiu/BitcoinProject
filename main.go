package main

import (
	_ "BitcoinProject/routers"
	"BitcoinProject/utils"
	"fmt"
)

func main() {
	//设置静态资源路径
	//beego.SetStaticPath("/js","./static/js")
	//beego.SetStaticPath("/css","./static/css")
	//beego.SetStaticPath("/img","./static/img")
	//
	//beego.Run()
	//blockheaders,err := utils.GetBlockHeader("00000000c3dd0ea58aa5b0fa0c520e1fbdad693b0d8918cf715186c914472042")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(blockheaders.Chainwork)
	//MemPoolInfo,err := utils.GetMemPoolInfo()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Printf("%.8f",MemPoolInfo.Bytes)
	TxoutSetInfo,err := utils.GetTxoutSetInfo()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(TxoutSetInfo.Height)
}
