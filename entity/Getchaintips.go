package entity

type Getchaintips struct {
	Height    float64 `json:"height"`
	Hash      string `json:"hash"`
	Hranchlen float64 `json:"hranchlen"`
	Htatus    string `json:"htatus"`
}
