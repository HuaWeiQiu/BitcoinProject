package entity

type RPCRequest struct {
	Id      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []interface{} `json:"params"`
}
