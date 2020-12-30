package utils

import (
	"bytes"
	"encoding/gob"
)

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	gob.Register(map[string]interface{}{})
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
