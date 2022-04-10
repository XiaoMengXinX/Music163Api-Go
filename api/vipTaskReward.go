package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// VipTaskRewardAPI 领取 VIP 任务成长值奖励 API
const VipTaskRewardAPI = "/api/vipnewcenter/app/level/task/reward/get"

// vipTaskRewardReqJson 领取 VIP 任务成长值请求参数
type vipTaskRewardReqJson struct {
	TaskIds string `json:"taskIds"`
}

// CreateVipTaskRewardReqJson 创建 领取 VIP 任务成长值 请求json
func CreateVipTaskRewardReqJson(taskID string) string {
	reqBody := vipTaskRewardReqJson{TaskIds: taskID}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// VipTaskReward 领取 VIP 任务成长值奖励
func VipTaskReward(data utils.RequestData, taskID string) (result types.VipTaskRewardData, err error) {
	var options utils.EapiOption
	options.Path = VipTaskRewardAPI
	options.Url = "https://music.163.com/eapi/vipnewcenter/app/level/task/reward/get"
	reqBodyJson := CreateVipTaskRewardReqJson(taskID)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
