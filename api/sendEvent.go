package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/google/uuid"
	"strings"
)

// SendEventAPI 发送动态 API
const SendEventAPI = "/api/share/friends/resource"

// SendEventReq SendEvent API 的 body json
type SendEventReq struct {
	Msg        string `json:"msg"`
	Type       string `json:"type"`
	UUID       string `json:"uuid"`
	Pics       string `json:"pics"`
	AddComment string `json:"addComment"`
	Header     string `json:"header"`
	ER         string `json:"e_r"`
}

// CreateTextEventReqJson 创建请求 body json
func CreateTextEventReqJson(text string) string {
	UUID := uuid.New()
	shareConfig := SendEventReq{
		Msg:        text,
		Type:       "noresource",
		UUID:       strings.Replace(UUID.String(), "-", "", -1),
		Pics:       "[]",
		AddComment: "false",
		Header:     "{}",
		ER:         "true",
	}
	reqBodyJson, _ := json.Marshal(shareConfig)
	return string(reqBodyJson)
}

// SendTextEvent 发送文本动态
func SendTextEvent(data utils.RequestData, text string) (result types.SendEventData, err error) {
	var options utils.EapiOption
	options.Path = SendEventAPI
	options.Url = "https://music.163.com/eapi/share/friends/resource"
	options.Json = CreateTextEventReqJson(text)
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
