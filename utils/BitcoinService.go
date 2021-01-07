package utils

import (
	"fmt"
	"github.com/mapstructure"
	"net/http"
	"strings"
	"BitcoinProject/entity"

)

/*
1、获取最新区块的hash
*/
func GetBestBlockHash() string {
	jsons := PrepareJSON("getbestblockhash", nil)
	jsonbody := strings.NewReader(jsons)
	rpcResult, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return ""
	}
	if rpcResult.Code == http.StatusOK {
		return rpcResult.Data.Result.(string)
	}
	return ""
}

/*
2、根据特定高度获取对应区块hash
*/
func GetBlockHash(height int) (string, error) {
	jsons := PrepareJSON("getblockhash", []interface{}{height})
	jsonbody := strings.NewReader(jsons)
	rpcResult, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return "", nil
	}
	if rpcResult.Code == 200 {
		return rpcResult.Data.Result.(string), nil
		//resultByte := []byte(resultStr)
		//err = json.Unmarshal([]byte(resultStr),&entity.Balance{})
		//if err!=nil{
		//	return &entity.Balance{},err
		//}
		//return &entity.Balance{},nil
	}
	return "", nil
}

/*
 *3、根据比特币地址获取地址对应的信息
 */
func GetBalances() (*entity.Balance, error) {
	jsons := PrepareJSON("getbalances", nil)
	jsonbody:=strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		fmt.Println(err.Error())
		return nil,nil
	}
	if result.Code == http.StatusOK {
		fmt.Println(result.Code)
		resultStr := result.Data.Result
		var balance *entity.Balance
		err:=mapstructure.WeakDecode(resultStr,&balance)
		if err != nil {
			fmt.Println(err.Error())
			return nil,err
		}
		return balance,err
	}
	return nil,err
}
/*
 *4、获取全部链头区块
 */
func GetChaintips() ([]*entity.Chaintips,error) {
	jsons := PrepareJSON("getchaintips",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil,nil
	}
	if result.Code == http.StatusOK {
		resultStr:=result.Data.Result
		//fmt.Println(result.Data.Result)
		var chaintips []*entity.Chaintips
		err:=mapstructure.WeakDecode(resultStr,&chaintips)
		if err != nil {
			fmt.Println(err.Error())
			return nil,nil
		}
		return chaintips,nil
	}
	return nil,nil
}
/*
 *5、获取区块的难度
 */
func GetDifficult() float64 {
	jsons := PrepareJSON("getdifficulty",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return -1
	}
	if result.Code == http.StatusOK {
		return result.Data.Result.(float64)
	}
	return -1
}

/*
 *6、获取比特币内存信息
 *@author  xcp
 */
func GetMemoryInfo() (*entity.MemoryInfo, error) {
	jsons := PrepareJSON("getmemoryinfo", nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil, nil
	}
	if result.Code == http.StatusOK {
		resultStr := result.Data.Result
		var memoryinfo entity.MemoryInfo
		err:=mapstructure.WeakDecode(resultStr,&memoryinfo)
		return &memoryinfo,err
	}
	return nil, nil
}

/*
 *7、获取区块链当前状态信息
 *@author yexin
 */
func GetBlockChainInfo() (*entity.Blockchaininfo,error) {
	jsons := PrepareJSON("getblockchaininfo",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil,nil
	}
	if result.Code == http.StatusOK {
		resultStr:=result.Data.Result
		var blockchaininfo entity.Blockchaininfo
		err:=mapstructure.WeakDecode(resultStr,&blockchaininfo)
		if err != nil {
			fmt.Println(err.Error())
			return nil,nil
		}
		return &blockchaininfo,nil
	}
	return nil,nil
}
/*
 *获取区块数量
 *8、getblockcount调用返回本地最优链中的区块数量
 *@author yexin
 */
func GetBlockCount() float64  {
	jsons := PrepareJSON("getblockcount",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return -1
	}
	if result.Code == http.StatusOK {
		return result.Data.Result.(float64)
	}
	return -1
}

/*
 *9、获取比特币远程调用(rpc)信息
 *@author  xcp
 */
func GetRpcInfo() (*entity.RPCInfo, error) {
	jsons := PrepareJSON("getrpcinfo", nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil, nil
	}
	if result.Code == http.StatusOK {
		resultStr := result.Data.Result
		var rpcinfo entity.RPCInfo
		err:=mapstructure.WeakDecode(resultStr,&rpcinfo)
		return &rpcinfo,err
	}
	return nil,nil
}

/*
 *10、获取日志记录
 *@author  xcp
 */
func Logging() (*entity.LOGGing, error) {
	jsons := PrepareJSON("logging", nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil, nil
	}
	if result.Code == http.StatusOK {
		resultStr := result.Data.Result
		var logging entity.LOGGing
		err:=mapstructure.WeakDecode(resultStr,&logging)
		return &logging,err
	}
	return nil,nil
}

/*
 *11、获取正常运行时间
 *@author  xcp
 */
func Uptime() float64 {
	jsons := PrepareJSON("uptime",nil)
	jsonbody:=strings.NewReader(jsons)
	result,err:=DoPost(RPCURL,ReqHeader(),jsonbody)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	if result.Code == http.StatusOK {
		return result.Data.Result.(float64)
	}
	return -1
}
/*
 *12、获取指定区块头
 */
func GetBlockHeader(hash string) ( *entity.Blockheader,error) {
	jsons := PrepareJSON("getblockheader",[]interface{}{hash})
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil,nil
	}
	if result.Code == http.StatusOK {
		resultStr:=result.Data.Result
		var blockheader entity.Blockheader
		err:=mapstructure.WeakDecode(resultStr,&blockheader)
		if err != nil {
			return nil,nil
		}
		return &blockheader,nil
	}
	return nil,nil
}
/*
 *13、getmempoolinfo
 */
func GetMemPoolInfo() (*entity.Mempoolinfo,error) {
	jsons := PrepareJSON("getmempoolinfo",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil,nil
	}
	if result.Code == http.StatusOK {
		resultStr:=result.Data.Result
		var mempoolinfo entity.Mempoolinfo
		err:=mapstructure.WeakDecode(resultStr,&mempoolinfo)
		if err != nil {
			fmt.Println(err.Error())
			return nil,nil
		}
		return &mempoolinfo,nil
	}
	return nil,nil
}
/*
 *14、获取交易输出的信息
 */
func GetTxoutSetInfo() (*entity.Txoutsetinfo,error)  {
	jsons := PrepareJSON("gettxoutsetinfo",nil)
	jsonbody := strings.NewReader(jsons)
	result,err := DoPost(RPCURL,ReqHeader(),jsonbody)
	if err != nil {
		return nil,nil
	}
	if result.Code ==http.StatusOK {
		resultStr :=result.Data.Result
		var txoutsetinfo entity.Txoutsetinfo
		err := mapstructure.WeakDecode(resultStr,&txoutsetinfo)
		if err != nil {
			fmt.Println(err.Error())
			return nil,nil
		}
		return &txoutsetinfo,nil
	}
	return nil,nil
}
/*
 *15、根据区块的hash获取区块的信息
*/
func GetBlock(hash string) (*entity.Block,error) {
	json := PrepareJSON("getblock",[]interface{}{hash})
	fmt.Println(json)
	result,err := DoPost(RPCURL,ReqHeader(),strings.NewReader(json))
	if err != nil {
		return nil,nil
	}
	if result.Code==http.StatusOK{
		resultStr := result.Data.Result
		var block entity.Block
		err= mapstructure.WeakDecode(resultStr,&block)
		if err!=nil {
			fmt.Println(err.Error())
			return nil,err
		}
		return &block,nil
	}
	return nil,nil
}
/*
**16、getnewaddress 生成一个新的比特币地址
*/
func GetNewAddress(label string,addresstype string)string{
	//addresstype:=entity.AddressTypes(addtype)
	json:=PrepareJSON("getnewaddress",[]interface{}{label,addresstype})
	body:=strings.NewReader(json)
	result,err:=DoPost(RPCURL,ReqHeader(),body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	if result.Code==http.StatusOK {
		return result.Data.Result.(string)
	}
	return ""
}
