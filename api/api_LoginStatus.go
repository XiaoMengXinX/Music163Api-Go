package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// LoginStatus 获取账号登录状态 API
const LoginStatus = "/api/w/nuser/account/get"

// GetLoginStatus 获取账号登录状态
func GetLoginStatus(data utils.RequestData) (result types.LoginStatusData, err error) {
	var options utils.EapiOption
	options.Path = LoginStatus
	options.Url = "https://music.163.com/eapi/w/nuser/account/get"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
