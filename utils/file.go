package utils

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// ReadFile 读取文件
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Errorln(err)
		}
	}(file)
	fileContent, err := ioutil.ReadAll(file)
	return fileContent, err
}

// DetectFileType 检测文件类型
func DetectFileType(file []byte) (fileType, fileSubtype string) {
	fileMime := http.DetectContentType(file)
	fileTypes := strings.Split(fileMime, "/")
	if len(fileTypes) == 2 {
		fileType = fileTypes[0]
		fileSubtype = fileTypes[1]
	}
	return
}
