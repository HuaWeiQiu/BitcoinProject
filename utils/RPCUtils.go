package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testbitcoin/entity"
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
func DoPost(url string, header map[string]string, body io.Reader) []byte {
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
	//判断响应是否成功
	var respbyte []byte
	if resp.StatusCode == http.StatusOK {
		respbyte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	return respbyte
}
