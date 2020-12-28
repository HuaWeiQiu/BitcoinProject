package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	//"testbitcoin/entity"
	"BitcoinProject/entity"
)

/*
获取最新区块的hash
*/
func GetBestBlockHash() string {
	jsons := PrepareJSON("getbestblockhash", nil)
	jsonbody := strings.NewReader(jsons)
	rpcResult := DoPost(RPCURL, ReqHeader(), jsonbody)
	var results entity.Result
	err := json.Unmarshal(rpcResult, &results)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return results.Result.(string)
}

/*
根据特定高度获取对应区块hash*/
//func GetBlockHash(height int) (interface{}, error) {
//	jsons := PrepareJSON("getblockhash", []interface{}{height})
//	jsonbody := strings.NewReader(jsons)
//	rpcResult, err := DoPost(RPCURL, ReqHeader(), jsonbody)
//	if err != nil {
//		return nil, nil
//	}
//	if rpcResult.Code == 200 {
//		return rpcResult.Data.Result, nil
//		//resultByte := []byte(resultStr)
//		//err = json.Unmarshal([]byte(resultStr),&entity.Balance{})
//		//if err!=nil{
//		//	return &entity.Balance{},err
//		//}
//		//return &entity.Balance{},nil
//	}
//	return nil, nil
//}

/*
根据比特币地址获取地址对应的信息
*/
func GetBalances() (*entity.Balance, error) {
	jsons := PrepareJSON("getbalances", nil)
	jsonbody := strings.NewReader(jsons)
	result := DoPost(RPCURL, ReqHeader(), jsonbody)
	//resultStr := result.Result
	//[]byte(resultStr)
	//var balance entity.Balance
	//err := mapstructure.WeakDecode(resultStr, &balance)
	//fmt.Println(result)
	//str := fmt.Sprintf("%v", resultStr)
	//fmt.Println("000000", str)
	//str1 := strings.ReplaceAll(str, "map[", "{")
	//str1 = strings.ReplaceAll(str1, "]", "}")
	//fmt.Println(str1)
	//fmt.Println(resultStr)
	//fmt.Println(resultStr)
	//Strbuffer := resultStr.([]byte)
	//var resultBytes = resultStr.([]byte)
	//Bytes, err := GetBytes(str1)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//str := hex.EncodeToString(Bytes)
	//fmt.Println(str)
	//fmt.Println(Bytes)
	//Bytes = bytes.TrimPrefix(Bytes, []byte("<"))
	//Bytes = bytes.TrimPrefix(Bytes, []byte("\f"))
	var balance entity.Balance
	err := json.Unmarshal(result, &balance)
	if err != nil {
		return &entity.Balance{}, err
	}
	return &balance, nil
}
func GetDifficulty() float64 {
	jsons := PrepareJSON("getdifficulty", nil)
	jsonbody := strings.NewReader(jsons)
	rpcResult := DoPost(RPCURL, ReqHeader(), jsonbody)
	var results entity.Result
	err := json.Unmarshal(rpcResult, &results)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return results.Result.(float64)
}
func Getchaintips() *entity.Getchaintips {
	jsons := PrepareJSON("getchaintips", nil)
	jsonbody := strings.NewReader(jsons)
	rpcResult := DoPost(RPCUSER, ReqHeader(), jsonbody)
	var results entity.Getchaintips
	fmt.Println(results)
	err := json.Unmarshal(rpcResult, &results)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &results
}

func Getblockhash(height int) string {
	jsons := PrepareJSON("getblockhash", []interface{}{height})
	jsonbody := strings.NewReader(jsons)
	rpcResult := DoPost(RPCUSER, ReqHeader(), jsonbody)
	var results entity.Result
	err := json.Unmarshal(rpcResult, &results)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return results.Result.(string)

}
