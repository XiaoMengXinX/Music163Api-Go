package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// VipTaskRewardAllAPI 领取所有 VIP 任务成长值奖励 API
const VipTaskRewardAllAPI = "/api/vipnewcenter/app/level/task/reward/getall"

// VipTaskRewardAll 领取所有 VIP 任务成长值奖励
func VipTaskRewardAll(data utils.RequestData) (result types.VipTaskRewardData, err error) {
	var options utils.EapiOption
	options.Path = VipTaskRewardAllAPI
	options.Url = "https://music.163.com/eapi/vipnewcenter/app/level/task/reward/getall"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
