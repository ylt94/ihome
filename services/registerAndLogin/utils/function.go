package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptionByMD5(str string) string {
	M := md5.New()
	M.Write([]byte(str))
	return hex.EncodeToString(M.Sum(nil))
}

func CreateTokenByJWT(phone string) string {
	//TODO JWT 生成登录TOKEN
	return ""
}
