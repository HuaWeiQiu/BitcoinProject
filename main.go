package main

import (
	_ "BitcoinProject/routers"
	"BitcoinProject/utils"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//getchaintips :=utils.Getchaintips()
	//fmt.Println(getchaintips)
	balances, err := utils.GetBalances()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(balances.Mine)
	bestblockhash := utils.GetBestBlockHash()
	fmt.Println(bestblockhash)
	hash := utils.Getblockhash(0)
	fmt.Println("===================",hash)
	beego.Run()
}
