package utils

import (
	"fmt"
	"github.com/mapstructure"
	"net/http"
	"strings"
	"BitcoinProject/entity"
)

/*
获取最新区块的hash
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
根据特定高度获取对应区块hash
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
 *根据比特币地址获取地址对应的信息
 */
func GetBalances() (*entity.Balance, error) {
	jsons := PrepareJSON("getbalances",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil, nil
	}
	//results :=result.Data.Result.([]byte)
	//fmt.Println(results)
	if result.Code == http.StatusOK {
		fmt.Println(result.Code)
		resultStr := result.Data.Result
		//[]byte(resultStr)
		var balance entity.Balance
		err:=mapstructure.WeakDecode(resultStr,&balance)
		////fmt.Println(result)
		//str := fmt.Sprintf("%v", resultStr)
		//fmt.Println("000000",str)
		//
		//str1 :=strings.ReplaceAll(str,"map[","{")
		//str1 =strings.ReplaceAll(str1,"]","}")
		//fmt.Println(str1)
		//
		////fmt.Println(resultStr)
		////fmt.Println(resultStr)
		////Strbuffer := resultStr.([]byte)
		////var resultBytes = resultStr.([]byte)
		//Bytes,err := GetBytes(str1)
		//if err != nil {
		//	fmt.Println(err.Error())
		//}
		////str := hex.EncodeToString(Bytes)
		////fmt.Println(str)
		//fmt.Println(Bytes)
		//
		//Bytes = bytes.TrimPrefix(Bytes,[]byte("<"))
		//Bytes = bytes.TrimPrefix(Bytes,[]byte("\f"))
		//var balance  entity.Balance
		//err = json.Unmarshal(Bytes, &balance)
		//if err != nil {
		//	return &entity.Balance{}, err
		//}
		//result.Data.Result = balance
		//return &entity.Balance{}, nil
		return &balance,err
	}
	return nil, nil
	//return nil,nil
}
/*
 *获取区块的难度
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
 *获取全部链头区块
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
		//fmt.Println(chaintips[0])
		return chaintips,nil
	}
	return nil,nil
}
/*
 *获取区块链当前状态信息
 *@author yexin
 */
func GetBlockChainInfo() (*entity.Blockchaininfo,error) {
	jsons := PrepareJSON("getblockchaininfo",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil,nil
	}
	//fmt.Println(result.Code)
	//fmt.Println(result.Data.Result)
	if result.Code == http.StatusOK {
		resultStr:=result.Data.Result
		var blockchaininfo entity.Blockchaininfo
		err:=mapstructure.WeakDecode(resultStr,&blockchaininfo)
		if err != nil {
			fmt.Println(err.Error())
			return nil,nil
		}
		//fmt.Println(blockchaininfo)
		return &blockchaininfo,nil
	}
	return nil,nil
}
/*
 *获取区块数量
 *getblockcount调用返回本地最优链中的区块数量
 *@author yexin
 */
func GitBlockCount() float64  {
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
 *获取指定区块头
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
 *getmempoolinfo
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
 *
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