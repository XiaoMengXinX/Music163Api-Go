package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserStatusDetailAPI 获取用户状态详情
const UserStatusDetailAPI = "/api/social/user/status/detail"

// UserStatusEditAPI 编辑用户状态
const UserStatusEditAPI = "/api/social/user/status/edit"

// UserStatusDetailReqJson 获取用户状态详情参数
type UserStatusDetailReqJson struct {
	VisitorId int `json:"visitorId"`
}

// UserStatusEditReqJson 编辑用户状态参数
type UserStatusEditReqJson struct {
	Content string `json:"content"`
}

type userStatusContent struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// CreateUserStatusDetailReqJson 创建请求 json
func CreateUserStatusDetailReqJson(visitorId int) string {
	reqBodyJson, _ := json.Marshal(UserStatusDetailReqJson{
		VisitorId: visitorId,
	})
	return string(reqBodyJson)
}

// CreateUserStatusEditReqJson 创建请求 json
func CreateUserStatusEditReqJson(context string) string {
	s, _ := json.Marshal(userStatusContent{
		Type:    "TEXT",
		Content: context,
	})
	reqBodyJson, _ := json.Marshal(UserStatusEditReqJson{
		Content: string(s),
	})
	return string(reqBodyJson)
}

// GetUserStatus 获取用户状态
func GetUserStatus(data utils.RequestData, userID int) (result types.UserStatusDetailData, err error) {
	var options utils.EapiOption
	options.Path = UserStatusDetailAPI
	options.Url = "https://music.163.com/eapi/social/user/status/detail"
	options.Json = CreateUserStatusDetailReqJson(userID)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// SetUserStatus 设置用户状态
func SetUserStatus(data utils.RequestData, ctx string) (result types.UserStatusDetailData, err error) {
	var options utils.EapiOption
	options.Path = UserStatusEditAPI
	options.Url = "https://music.163.com/eapi/social/user/status/edit"
	options.Json = CreateUserStatusEditReqJson(ctx)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
