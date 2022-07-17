package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// CloudbeanNumAPI 获取音乐人云豆数量 API
const CloudbeanNumAPI = "/api/cloudbean/get"

// GetCloudbeanNum 获取音乐人云豆数量
func GetCloudbeanNum(data utils.RequestData) (result types.CloudBeanNumData, err error) {
	var options utils.EapiOption
	options.Path = CloudbeanNumAPI
	options.Url = "https://music.163.com/eapi/cloudbean/get"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
