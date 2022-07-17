package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserSignAPI 用户签到 API
const UserSignAPI = "/api/point/dailyTask"

// UserSignReqJson 签到参数
type userSignReqJson struct {
	Type int `json:"type"`
}

// CreateUserSignReqJson 创建 用户签到 请求json,
// signType 为签到类型 ,默认 0 ,其中0 为安卓端签到 ,1 为 web/PC 签到
func CreateUserSignReqJson(signType int) string {
	reqBody := userSignReqJson{}
	switch signType {
	case 0:
		reqBody.Type = 0
	case 1:
		reqBody.Type = 1
	default:
		reqBody.Type = 0
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// UserSign 用户签到, signType 为签到类型 ,默认 0 ,其中0 为安卓端签到 ,1 为 web/PC 签到
func UserSign(data utils.RequestData, signType int) (result types.UserSignData, err error) {
	var options utils.EapiOption
	options.Path = UserSignAPI
	options.Url = "https://music.163.com/eapi/point/dailyTask"
	reqBodyJson := CreateUserSignReqJson(signType)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
