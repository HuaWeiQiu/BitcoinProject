package entity

type RPCInfo struct {
	Active_commands Active_commands
	Logpath string
}

type Active_commands struct {
	Method string
	Duration int
}