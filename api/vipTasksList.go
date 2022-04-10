package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// VipTasksListAPI 获取 VIP 任务列表 API
const VipTasksListAPI = "/api/vipnewcenter/app/level/task/list"

// GetVipTasksList 获取 VIP 任务列表
func GetVipTasksList(data utils.RequestData) (result types.VipTaskslistDetailData, err error) {
	var options utils.EapiOption
	options.Path = VipTasksListAPI
	options.Url = "https://music.163.com/eapi/vipnewcenter/app/level/task/list"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
