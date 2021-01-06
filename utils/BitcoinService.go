package utils

import (
	"BitcoinProject/entity"
	"github.com/mapstructure"
	"net/http"
	"strings"
)

//Control

/*
 *获取比特币内存信息
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
	return nil,nil
}

/*
 *获取比特币远程调用(rpc)信息
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
 *获取日志记录
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
 *获取正常运行时间
 *@author  xcp
 */
func Uptime() int {
	jsons := PrepareJSON("uptime",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return -1
	}

	if result.Code == http.StatusOK {
		return result.Data.Result.(int)
	}
	return -1
}

//Network

/*
 *获取连接数
 *@author  xcp
 */
func GetConnectionCount() int {
	jsons := PrepareJSON("getconnectioncount",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return -1
	}

	if result.Code == http.StatusOK {
		return result.Data.Result.(int)
	}
	return -1
}

/*
 *获取净总数
 *@author  xcp
 */
func GetNetTotals() (*entity.Nettotals, error) {
	jsons := PrepareJSON("getnettotals", nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil, nil
	}
	if result.Code == http.StatusOK {
		resultStr := result.Data.Result
		var nettotals entity.Nettotals
		err:=mapstructure.WeakDecode(resultStr,&nettotals)
		return &nettotals,err
	}
	return nil,nil
}

/*
 *获取节点地址
 *@author  xcp
 */

func GetNodeAddresses() ([]*entity.Nodeaddresses, error) {
	jsons := PrepareJSON("getnodeaddresses", nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil, nil
	}
	if result.Code == http.StatusOK {
		resultStr := result.Data.Result
		var nodeaddresses []*entity.Nodeaddresses
		err:=mapstructure.WeakDecode(resultStr,&nodeaddresses)
		return nodeaddresses,err
	}
	return nil, nil
}

/*
 *获取同行的信息
 *@author  xcp
 */

func GetPeerInfo() []int {
	jsons := PrepareJSON("getpeerinfo",nil)
	jsonbody := strings.NewReader(jsons)
	result, err := DoPost(RPCURL, ReqHeader(), jsonbody)
	if err != nil {
		return nil
	}

	if result.Code == http.StatusOK {
		return result.Data.Result.([]int)
	}
	return nil
}
