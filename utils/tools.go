package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// RandHex 生成随机十六进制字符串
func RandHex(n int) []byte {
	rand.Seed(time.Now().Unix())
	if n <= 0 {
		return []byte{}
	}
	var need int
	if n&1 == 0 {
		need = n
	} else {
		need = n + 1
	}
	size := need / 2
	dst := make([]byte, need)
	src := dst[size:]
	if _, err := rand.Read(src[:]); err != nil {
		return []byte{}
	}
	hex.Encode(dst, src)
	return dst[:n]
}
