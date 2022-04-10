package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// DelEventAPI 删除动态 API
const DelEventAPI = "/api/event/delete"

// delEventReq DelEvent API 的 body json
type delEventReq struct {
	ID string `json:"id"`
}

// CreateDelEventReqJson 创建 删除动态 请求json
func CreateDelEventReqJson(eventID int) string {
	delEventConfig := delEventReq{
		ID: fmt.Sprintf("%d", eventID),
	}
	bodyJson, _ := json.Marshal(delEventConfig)
	return string(bodyJson)
}

// DelEvent 删除动态
func DelEvent(data utils.RequestData, eventID int) (result types.DelEventData, err error) {
	var options utils.EapiOption
	options.Path = DelEventAPI
	options.Url = "https://music.163.com/eapi/event/delete"
	options.Json = CreateDelEventReqJson(eventID)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
