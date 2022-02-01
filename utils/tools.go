package utils

import (
	"bytes"
	"encoding/hex"
	"image"
	// 解析 gif 图片信息
	_ "image/gif"
	// 解析 jpg 图片信息
	_ "image/jpeg"
	// 解析 png 图片信息
	_ "image/png"
	l "log"
	"math/rand"
	"time"
)

var log Logger = l.Default()

// Logger 日志接口
type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// SetLogger 设置日志记录器
func SetLogger(logger Logger) {
	log = logger
}

// ImageSize 获取图片尺寸
func ImageSize(file []byte) (width, height int, err error) {
	reader := bytes.NewReader(file)
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return
	}
	return img.Width, img.Height, err
}

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
