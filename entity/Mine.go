package entity

type Mine struct {
	Trusted           float64 `mapstructure:"trusted"`
	Untrusted_pending float64 `mapstructure:"untrusted_pending"`
	Immature          float64 `mapstructure:"Immature"`
	//Trusted           float64 `json:"trusted"`
	//Untrusted_pending float64 `json:"untrusted_pending"`
	//Immature          float64 `json:"immature"`
}
