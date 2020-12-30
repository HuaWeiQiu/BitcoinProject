package entity

type Mempoolinfo struct {
	Loaded        bool    `mapstructure:"loaded"`
	Size          float64 `mapstructure:"size"`
	Bytes         float64 `mapstructure:"bytes"`
	Usage         float64 `mapstructure:"usage"`
	Maxmempool    float64 `mapstructure:"maxmempool"`
	Mempoolminfee float64 `mapstructure:"mempoolminfee"`
	Minrelaytxfee float64 `mapstructure:"minrelaytxfee"`
}
