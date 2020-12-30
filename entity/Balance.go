package entity

type Balance struct {
	Mine Mine `mapstructure:"mine"`
	//Mine Mine `json:"mine"`
}
type Mine struct {
	Trusted           float64 `mapstructure:"trusted"`
	Untrusted_pending float64 `mapstructure:"untrusted_pending"`
	Immature          float64 `mapstructure:"Immature"`
	//Trusted           float64 `json:"trusted"`
	//Untrusted_pending float64 `json:"untrusted_pending"`
	//Immature          float64 `json:"immature"`
}

//type Balances struct {
//	Mine struct{
//		Trusted           float64
//		Untrusted_pending float64
//		Immature          float64
//	}
//}

//func Balancess()  {
//	balance := Balances{Mine: struct {
//		Trusted           float64
//		Untrusted_pending float64
//		Immature          float64
//	}{Trusted: 0, Untrusted_pending: 0, Immature: 0}}
//}
