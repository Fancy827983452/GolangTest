package algorithm

import (
	"crypto/md5"
	"encoding/hex"
)

// 生成32位MD5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}





