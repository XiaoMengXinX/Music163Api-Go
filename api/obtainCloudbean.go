package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// ObtainCloudbeanAPI 领取云豆 API
const ObtainCloudbeanAPI = "/api/nmusician/workbench/mission/reward/obtain/new"

// obtainCloudbeanReqJson 领取云豆参数
type obtainCloudbeanReqJson struct {
	UserMissionId int `json:"userMissionId"`
	Period        int `json:"period"`
}

// CreateObtainCloudbeanReqJson 创建 领取云豆 请求json
func CreateObtainCloudbeanReqJson(userMissionID, period int) string {
	reqBody := obtainCloudbeanReqJson{
		UserMissionId: userMissionID,
		Period:        period,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// ObtainCloudbean 领取云豆, userMissionID, period 需通过 GetMusicianTasks 获取
func ObtainCloudbean(data utils.RequestData, userMissionID, period int) (result types.ObtainCloudebeanData, err error) {
	var options utils.EapiOption
	options.Path = ObtainCloudbeanAPI
	options.Url = "https://music.163.com/eapi/nmusician/workbench/mission/reward/obtain/new"
	reqBodyJson := CreateObtainCloudbeanReqJson(userMissionID, period)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
