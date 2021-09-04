package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SendMsgAPI 发送私信 API
const SendMsgAPI = "/api/msg/private/send"

// SendMsgReq 发送私信 API 的 body json
type SendMsgReq struct {
	UserIds string `json:"userIds"`
	Msg     string `json:"msg"`
	Type    string `json:"type"`
	Header  string `json:"header"`
	ER      string `json:"e_r"`
}

// CreateTextMsgReqJson 创建请求 body json
func CreateTextMsgReqJson(userIDs []int, msg string) string {
	var userID []int
	for i := 0; i < len(userIDs); i++ {
		userID = append(userID, userIDs[i])
	}
	userIDJson, _ := json.Marshal(userID)
	reqBody := SendMsgReq{
		UserIds: string(userIDJson),
		Msg:     msg,
		Header:  "{}",
		Type:    "text",
		ER:      "true",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// SendTextMsg 发送文本私信
func SendTextMsg(data utils.RequestData, userIDs []int, text string) (result types.SendMsgData, err error) {
	var options utils.EapiOption
	options.Path = SendMsgAPI
	options.Url = "https://music.163.com/eapi/msg/private/send"
	options.Json = CreateTextMsgReqJson(userIDs, text)
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
