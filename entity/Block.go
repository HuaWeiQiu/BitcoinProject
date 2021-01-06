package entity

type Block struct {
	Hash          string   `mapstructure:"hash"`
	Confirmations int      `mapstructure:"confirmations"`
	Strippedsize  int      `mapstructure:"strippedsize"`
	Weight        int      `mapstructure:"weight"`
	Height        int      `mapstructure:"height"`
	Version       int      `mapstructure:"version"`
	VersionHex    string   `mapstructure:"versionHex"`
	Merkleroot    string   `mapstructure:"merkleroot"`
	Tx            []string `mapstructure:"tx"`
	Time          int64    `mapstructure:"time"`
	Mediantime    int64    `mapstructure:"mediantime"`
	Nonce         int64    `mapstructure:"nonce"`
	Bits          string   `mapstructure:"bits"`
	Difficulty    int      `mapstructure:"difficulty"`
	Chainwork     string   `mapstructure:"chainwork"`
	NTx           int      `mapstructure:"nTx"`
}
