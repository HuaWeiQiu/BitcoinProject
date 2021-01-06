package models

type Command struct {
	Method string `form:"method"`
	Params1 string `form:"params1"`
	Params2 string `form:"params2"`
}

