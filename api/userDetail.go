package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserDetailAPI 用户详细信息 API
const UserDetailAPI = "/api/v1/user/detail"

// GetUserDetail 获取用户详细信息
func GetUserDetail(data utils.RequestData, userID int) (result types.UserDetailData, err error) {
	var options utils.EapiOption
	options.Path = fmt.Sprintf("%s/%d", UserDetailAPI, userID)
	options.Url = "https://music.163.com/eapi/v1/user/detail"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
