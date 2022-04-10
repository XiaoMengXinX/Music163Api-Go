package api

import (
	"encoding/json"
	"net/http"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// QrLoginAPI 检查 QR 登录状态 API
const QrLoginAPI = "/api/login/qrcode/client/login"

// checkQrLoginReq CheckQrLogin API 的 body json
type checkQrLoginReq struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

// CreateCheckQrLoginJson 创建 检查Qr登录状态 请求json
func CreateCheckQrLoginJson(key string) string {
	reqBody := checkQrLoginReq{
		Key:  key,
		Type: "3",
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
	resBody, header, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, header, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, header, err
}
