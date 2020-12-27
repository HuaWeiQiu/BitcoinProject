package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/*
对一个字符串进行MD5哈希计算，并返回hash值
*/
func Md5Hash(data string) string {
	hashMd5 := md5.New()
	hashMd5.Write([]byte(data))
	databytes := hashMd5.Sum(nil)
	return hex.EncodeToString(databytes)

}
func Md5HashFile(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashByte := md5Hash.Sum(nil)
	return hex.EncodeToString(hashByte), nil
}
func Sha256Hash(data []byte) []byte {
	//1、对block字段进行拼接

	//2、对拼接后的数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(data))
	return sha256Hash.Sum(nil)
}
func Sha256FileHash(reader io.Reader) (string,error) {
	//1、对block字段进行拼接

	//2、对拼接后的数据进行sha256
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	sha256Hash := sha256.New()
	sha256Hash.Write(bytes)
	return hex.EncodeToString(sha256Hash.Sum(nil)), nil
}
