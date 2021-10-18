package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"net/http"
)

// QrLoginAPI 检查 QR 登录状态 API
const QrLoginAPI = "/api/login/qrcode/client/login"

// CheckQrLoginReq CheckQrLogin API 的 body json
type CheckQrLoginReq struct {
	Key    string `json:"key"`
	Type   string `json:"type"`
	ER     bool   `json:"e_r"`
	Header string `json:"header"`
}

// CreateCheckQrLoginJson 创建请求 body json
func CreateCheckQrLoginJson(key string) string {
	reqBody := CheckQrLoginReq{
		Key:    key,
		Type:   "3",
		ER:     true,
		Header: "{}",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// CheckQrLogin 检查 QR 登录状态
func CheckQrLogin(data utils.RequestData, key string) (result types.QrCheckData, header http.Header, err error) {
	var options utils.EapiOption
	options.Path = QrLoginAPI
	options.Url = "https://music.163.com/eapi/login/qrcode/client/login"
	options.Json = CreateCheckQrLoginJson(key)
	resBody, header, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, header, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, header, err
}
