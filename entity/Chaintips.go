package entity

type Chaintips struct {
	Height int `mapstructure:"height"`
	Hash string `mapstructure:"hash"`
	Branchlen int `mapstructure:"branchlen"`
	Status string `mapstructure:"status"`
}
