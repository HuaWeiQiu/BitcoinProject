package entity

type Blockchaininfo struct {
	Chain                string    `mapstructure:"chain"`
	Blocks               float64   `mapstructure:"blocks"`
	Headers              float64   `mapstructure:"headers"`
	Bestblockhash        string    `mapstructure:"bestblockhash"`
	Difficulty           float64   `mapstructure:"difficulty"`
	Mediantime           float64   `mapstructure:"mediantime"`
	Verificationprogress float64   `mapstructure:"verificationprogress"`
	Initialblockdownload bool      `mapstructure:"initialblockdownload"`
	Chainwork            string    `mapstructure:"chainwork"`
	Size_on_disk         float64   `mapstructure:"size_on_disk"`
	Pruned               bool      `mapstructure:"pruned"`
	Softforks            Softforks `mapstructure:"softforks"`
	Warnings             string    `mapstructure:"warnings"`
}
type Softforks struct {
	Bip34  Types `mapstructure:"bip34"`
	Bip66  Types `mapstructure:"bip34"`
	Bip65  Types `mapstructure:"bip65"`
	Csv    Types `mapstructure:"csv"`
	Segwit Types `mapstructure:"segwit"`
}

type Types struct {
	Type   string  `mapstructure:"type"`
	Active bool    `mapstructure:"active"`
	Height float64 `mapstructure:"height"`
}
