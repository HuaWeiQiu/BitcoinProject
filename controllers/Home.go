package controllers

import (
	"BitcoinProject/models"
	"BitcoinProject/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type HoemController struct {
	beego.Controller
}

func (h *HoemController) Get() {
	h.TplName = "home.html"
}

func (h *HoemController) Post() {
	//bestblockhash:=c.GetString("getbestblockhash")
	//var command models.Command
	//err:=c.ParseForm(&command)
	//if err != nil {
	//	c.Ctx.WriteString("获取方法参数失败,请重试")
	//	return
	//}
	//body:=utils.PrepareJSON(command.Method,command.Params)
	//result,err:=utils.DoPost(utils.RPCURL,utils.ReqHeader(),strings.NewReader(body))
	//if err != nil {
	//	c.Ctx.WriteString("获取数据失败,请重试！")
	//	return
	//}
	//code := result.Code
	//resultStr := result.Data.Result
	//if code==http.StatusOK {
	//	c.Data["result"]=resultStr
	//}else {
	//	return
	//}
	fmt.Println("是否为ajax请求", h.IsAjax())
	var command models.Command
	data := h.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//err:=c.ParseForm(&command);
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	fmt.Println(command)
	method := command.Method
	params1 := command.Params1
	params2:=command.Params2

	//method:=c.GetString("method")
	//params:=c.GetString("params")
	//command:=&command{
	//	method:method,
	//	params:params,
	//}

	//fmt.Println("获取的参数为",method,params)
	//err:=json.Unmarshal(c.Ctx.Input.RequestBody,command)
	fmt.Println("获取的参数为", command)
	//c.TplName="result.html"
	result := make(map[string]interface{})
	result["getbestblockhash"] = utils.GetBestBlockHash()
	result["getblockcount"] = utils.GetBlockCount()
	height, _ := strconv.Atoi(params1)
	result["getblockhash"], _ = utils.GetBlockHash(height)
	result["getbalances"], _ = utils.GetBalances()
	result["getblock"], _ = utils.GetBlock(params1)
	result["getchaintips"],_=utils.GetChaintips()
	result["getdifficulty"]=utils.GetDifficult()
	result["getmemoryinfo"],_=utils.GetMemoryInfo()
	result["getblockchaininfo"],_=utils.GetBlockChainInfo()
	result["getrpcinfo"],_=utils.GetRpcInfo()
	result["logging"],_=utils.Logging()
	result["uptime"]=utils.Uptime()
	result["getblockheader"],_=utils.GetBlockHeader(params1)
	result["getmempoolinfo"],_=utils.GetMemPoolInfo()
	result["gettxoutsetinfo"],_=utils.GetTxoutSetInfo()
	result["getnewaddress"]=utils.GetNewAddress(params1,params2)
	for key, value := range result {
		if method == key {
			//fmt.Println(value)
			//c.Data["Result"]=value
			results, err := json.Marshal(value)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			h.Data["json"] = string(results)
			h.ServeJSON()
		}
		continue
	}
}
