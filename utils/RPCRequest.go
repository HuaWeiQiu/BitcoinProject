package utils

type Json struct {
	Id      int64         `json:"id"`
	Method  string        `json:"method"`
	Jsonrpc string        `json:"jsonrpc"`
	Params  []interface{} `json:"params"`
}
