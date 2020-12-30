package utils

/*
请求头
*/
func ReqHeader() map[string]string {
	header := make(map[string]string)
	header["Encoding"] = "UTF-8"
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Basic " + Base64ToStr(RPCUSER+":"+RPCPASSWORD)
	return header
}
