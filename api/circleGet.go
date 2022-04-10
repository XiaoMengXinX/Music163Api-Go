package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// GetCircleAPI 获取云圈动态API
const GetCircleAPI = `/api/circle/get`

// GetCircleReq 获取云圈动态API 的 body json
type GetCircleReq struct {
	CircleID string `json:"circleId"`
}

// CreateGetCircleReqJson 创建获取云圈动态API 的 body json
func CreateGetCircleReqJson(circleID string) string {
	reqBody := GetCircleReq{
		CircleID: circleID,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetCircle 获取云圈动态
func GetCircle(data utils.RequestData, cricleID string) (result types.GetCircleData, err error) {
	var options utils.EapiOption
	options.Path = GetCircleAPI
	options.Url = "https://music.163.com/eapi/circle/get"
	options.Json = CreateGetCircleReqJson(cricleID)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
