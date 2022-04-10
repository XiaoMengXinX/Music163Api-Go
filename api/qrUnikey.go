package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// QrUnikeyAPI 获取 QR 登录的 unikey
const QrUnikeyAPI = "/api/login/qrcode/unikey"

// qrUniKeyReq QrUnikey API 的 body json
type qrUniKeyReq struct {
	Type string `json:"type"`
}

// CreateGetQrUnikeyJson 创建 获取Qr登录key 请求json
func CreateGetQrUnikeyJson() string {
	reqBody := qrUniKeyReq{
		Type: "3",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetQrUnikey 获取 QR 登录的 unikey
func GetQrUnikey(data utils.RequestData) (result types.QrUnikeyData, err error) {
	var options utils.EapiOption
	options.Path = QrUnikeyAPI
	options.Url = "https://music.163.com/eapi/login/qrcode/unikey"
	options.Json = CreateGetQrUnikeyJson()
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
