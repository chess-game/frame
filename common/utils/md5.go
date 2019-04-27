package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/*
	md5加密
*/
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
