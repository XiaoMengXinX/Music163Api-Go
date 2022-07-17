package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// VipInfoAPI 获取 VPI 详细信息 API
const VipInfoAPI = "/api/music-vip-membership/client/vip/info"

// GetVipInfo 获取当前账号 VIP 信息
func GetVipInfo(data utils.RequestData) (result types.VipInfoData, err error) {
	var options utils.EapiOption
	options.Path = VipInfoAPI
	options.Url = "https://music.163.com/eapi/music-vip-membership/client/vip/info"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
