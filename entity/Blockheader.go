package entity

type Blockheader struct {
	Hash              string  `mapstructure:"hash"`
	Confirmations     float64 `mapstructure:"confirmations"`
	Height            float64 `mapstructure:"height"`
	Version           float64 `mapstructure:"version"`
	VersionHex        string  `mapstructure:"versionHex"`
	Merkleroot        string  `mapstructure:"merkleroot"`
	Time              float64 `mapstructure:"time"`
	Mediantime        float64 `mapstructure:"mediantime"`
	Nonce             float64 `mapstructure:"nonce"`
	Bits              string  `mapstructure:"bits"`
	Difficulty        float64 `mapstructure:"difficulty"`
	Chainwork         string  `mapstructure:"chainwork"`
	NTx               float64 `mapstructure:"nTx"`
	Previousblockhash string  `mapstructure:"previousblockhash"`
}
