package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// VipGrowrthpointAPI VIP 成长值信息 API
const VipGrowrthpointAPI = "/api/vipnewcenter/app/level/growhpoint/basic"

// GetVipGrowrthpoint 获取当前账号 VIP 成长值信息
func GetVipGrowrthpoint(data utils.RequestData) (result types.VipGrowthpointData, err error) {
	var options utils.EapiOption
	options.Path = VipGrowrthpointAPI
	options.Url = "https://music.163.com/eapi/vipnewcenter/app/level/growhpoint/basic"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
