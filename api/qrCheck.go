package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// QrLogin 检查 QR 登录状态 API
const QrLogin = "/api/login/qrcode/client/login"

// CheckQrLoginReq CheckQrLogin API 的 body json
type CheckQrLoginReq struct {
	Key    string `json:"key"`
	Type   string `json:"type"`
	ER     bool   `json:"e_r"`
	Header string `json:"header"`
}

// CreateCheckQrLoginJSON 创建请求 body json
func CreateCheckQrLoginJSON(key string) string {
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
func CheckQrLogin(data utils.RequestData, key string) (result types.QrCheckData, resHeader string, err error) {
	var options utils.EapiOption
	options.Path = QrLogin
	options.Url = "https://music.163.com/eapi/v3/song/detail"
	options.Json = CreateCheckQrLoginJSON(key)
	resBody, resHeader, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, resHeader, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, resHeader, err
}
