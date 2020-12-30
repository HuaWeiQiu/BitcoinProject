package utils

import (
	"BitcoinConnect/entity"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)
/*
准备一个json数据
*/
func PrepareJSON(method string, params []interface{}) string {
	rpcReq := entity.RPCRequest{
		Id:      time.Now().Unix(),
		Method:  method,
		Jsonrpc: RPCUCERSION,
		Params:  params,
	}
	reqBytes, err := json.Marshal(&rpcReq)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(reqBytes)
}
/**
执行post请求
 */
func DoPost(url string, header map[string]string, body io.Reader) (*entity.RPCResult, error) {
	client := &http.Client{}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err.Error())
	}
	//设置请求头
	if header != nil {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}
	//发起请求
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	code := resp.StatusCode
	//fmt.Println(code)
	rpcResult := entity.RPCResult{}
	if code == 200 {
		rpcResult.Code = code
		rpcResult.Msg = "请求成功"
		respbyte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(respbyte)
		var result entity.Result
		//反序列化
		err = json.Unmarshal(respbyte, &result)
		if err != nil {
			return &rpcResult, err
		}
		rpcResult.Data = result
		return &rpcResult, nil
	}
	rpcResult.Code = code
	rpcResult.Msg = "请求失败"
	return &rpcResult, nil
}
