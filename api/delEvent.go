package api

import (
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// DelEvent 删除动态 API
const DelEventAPI = "/api/event/delete"

// DelEventReq DelEvent API 的 body json
type DelEventReq struct {
	ID     string `json:"id"`
	Header string `json:"header"`
	ER     string `json:"e_r"`
}

// CreateDelEventReqJson 创建请求 body json
func CreateDelEventReqJson(eventID int) string {
	delEventConfig := DelEventReq{
		ID:     fmt.Sprintf("%d", eventID),
		Header: "{}",
		ER:     "true",
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
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
