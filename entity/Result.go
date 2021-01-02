package entity

type Result struct {
	Id     int64  `json:"id"`
	Error  error  `json:"error"`
	Result interface{} `json:"result"`
}
