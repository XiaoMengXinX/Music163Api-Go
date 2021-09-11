package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserSetting 用户设置 API
const UserSetting = "/api/user/setting"

// GetUserSetting 获取用户信息
func GetUserSetting(data utils.RequestData) (result types.UserSettingData, err error) {
	var options utils.EapiOption
	options.Path = UserSetting
	options.Url = "https://music.163.com/eapi/user/setting"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
