package main

import (
	_ "BitcoinProject/routers"
	"BitcoinProject/utils"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	bestblockhash:=utils.GetBestBlockHash()
	fmt.Println(bestblockhash)
	beego.Run()
}

