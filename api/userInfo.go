package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserInfoAPI 用户信息 API
const UserInfoAPI = "/api/v1/user/info"

// GetUserInfo 获取用户信息
func GetUserInfo(data utils.RequestData) (result types.UserInfoData, err error) {
	var options utils.EapiOption
	options.Path = UserInfoAPI
	options.Url = "https://music.163.com/eapi/v1/user/info"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
