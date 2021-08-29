package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// QrUnikey 获取 QR 登录的 unikey
const QrUnikey = "/api/login/qrcode/unikey"

// QrUniKeyReq QrUnikey API 的 body json
type QrUniKeyReq struct {
	Type   string `json:"type"`
	ER     bool   `json:"e_r"`
	Header string `json:"header"`
}

// CreateGetQrUnikeyJSON 创建请求 body json
func CreateGetQrUnikeyJSON() string {
	reqBody := QrUniKeyReq{
		Type:   "3",
		ER:     true,
		Header: "{}",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetQrUnikey 获取 QR 登录的 unikey
func GetQrUnikey(data utils.RequestData) (result types.QrUnikeyData, err error) {
	var options utils.EapiOption
	options.Path = QrUnikey
	options.Url = "https://music.163.com/eapi/v3/song/detail"
	options.Json = CreateGetQrUnikeyJSON()
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
