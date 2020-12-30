package entity

type Txoutsetinfo struct {
	Height            float64 `mapstructure:"height"`
	Bestblock         string  `mapstructure:"bestblock"`
	Transactions      float64 `mapstructure:"transactions"`
	Txouts            float64 `mapstructure:"txouts"`
	Bogosize          float64 `mapstructure:"bogosize"`
	Hash_serialized_2 string  `mapstructure:"hash_serialized_2"`
	Disk_size         float64 `mapstructure:"disk_size"`
	Total_amount      float64 `mapstructure:"total_amount"`
}
