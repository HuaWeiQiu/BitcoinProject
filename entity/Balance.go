package entity

type Balance struct {
	Mine Mine `mapstructure:"mine"`
	//Mine Mine `json:"mine"`
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
