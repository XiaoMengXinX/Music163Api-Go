package api

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"path"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// NosTokenAPI 获取 NosToken API （用于文件上传）
const NosTokenAPI = "/api/nos/token/alloc"

// nosTokenReq 获取 NosToken 的 body json
type nosTokenReq struct {
	Filename   string `json:"filename"`
	Local      string `json:"local"`
	NosProduct int    `json:"nos_product"`
	FileSize   int    `json:"fileSize"`
	Type       string `json:"type"`
	Md5        string `json:"md5"`
	Ext        string `json:"ext"`
}

// CreateNosTokenReqJson 创建 获取NosToken 请求json
func CreateNosTokenReqJson(filePath string) (string, []byte, error) {
	file, err := utils.ReadFile(filePath)
	if len(file) <= 32 {
		return "", file, fmt.Errorf("Empty file %s ", filePath)
	}
	fileType, _ := utils.DetectFileType(file[:32])
	size := len(file)
	md5str := fmt.Sprintf("%+x", md5.Sum(file))
	fileName := path.Base(filePath)
	fileExt := path.Ext(filePath)
	reqBody := nosTokenReq{
		Filename:   fileName,
		Local:      "false",
		NosProduct: 0,
		FileSize:   size,
		Md5:        md5str,
		Ext:        fileExt,
		Type:       fileType,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson), file, err
}

// GetNosToken 获取 NosToken
func GetNosToken(data utils.RequestData, filePath string) (result types.NosTokenData, file []byte, err error) {
	var options utils.EapiOption
	options.Path = NosTokenAPI
	options.Url = "https://music.163.com/eapi/nos/token/alloc"
	reqBodyJson, file, err := CreateNosTokenReqJson(filePath)
	if err != nil {
		return result, file, err
	}
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, file, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, file, err
}

// mlogNosTokenReq 获取 NosToken 的 body json
type mlogNosTokenReq struct {
	BizKey   string `json:"bizKey"`
	Filename string `json:"filename"`
	Bucket   string `json:"bucket"`
	Md5      string `json:"md5"`
	Type     string `json:"type"`
	FileSize int    `json:"fileSize"`
}

// CreateMlogNosTokenReqJson 创建 获取mlog的NosToken 请求json
func CreateMlogNosTokenReqJson(filePath string) (string, []byte, error) {
	file, err := utils.ReadFile(filePath)
	if len(file) <= 32 {
		return "", file, fmt.Errorf("Empty file: %s ", filePath)
	}
	fileType, _ := utils.DetectFileType(file[:32])
	size := len(file)
	md5str := fmt.Sprintf("%+x", md5.Sum(file))
	fileName := path.Base(filePath)
	reqBody := mlogNosTokenReq{
		BizKey:   string(utils.RandHex(8)),
		Filename: fileName,
		Bucket:   "yyimgs",
		Md5:      md5str,
		Type:     fileType,
		FileSize: size,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson), file, err
}

// GetMlogNosToken 获取用于 Mlog 的 NosToken
func GetMlogNosToken(data utils.RequestData, filePath string) (result types.NosTokenData, file []byte, err error) {
	var options utils.EapiOption
	options.Path = "/api/nos/token/whalealloc"
	options.Url = "https://music.163.com/eapi/nos/token/whalealloc"
	reqBodyJson, file, err := CreateMlogNosTokenReqJson(filePath)
	if err != nil {
		return result, file, err
	}
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, file, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	result.Type = 1
	return result, file, err
}
