package utils

import "encoding/base64"
/**
base64转String
 */
func Base64ToStr(msg string)string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}