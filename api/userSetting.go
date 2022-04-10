package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserSettingAPI 用户设置 API
const UserSettingAPI = "/api/user/setting"

// GetUserSetting 获取用户信息
func GetUserSetting(data utils.RequestData) (result types.UserSettingData, err error) {
	var options utils.EapiOption
	options.Path = UserSettingAPI
	options.Url = "https://music.163.com/eapi/user/setting"
	resBody, _, err := utils.ApiRequest(options, data)
	_ = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
