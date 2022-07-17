package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// ResourceShareAPI 分享资源 API
const ResourceShareAPI = "/api/share/friends/resource"

// resourceShareReq resourceShare API 的 body json
type resourceShareReq struct {
	ResourceID   int    `json:"id"`
	ResourceType string `json:"type"`
	Msg          string `json:"msg"`
}

// CreateResourceShareReqJson 创建 分享资源 请求json
func CreateResourceShareReqJson(resourceID int, resourceType string, msg string) string {
	if resourceType == "" {
		resourceType = "song"
	}
	reqBody := resourceShareReq{
		ResourceID:   resourceID,
		ResourceType: resourceType,
		Msg:          msg,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// ShareResource 分享资源
// 参数说明：
// ResourceID: 资源ID（歌曲，歌单，mv，电台，电台节目对应 id）
// ResourceType: 资源类型，默认歌曲 song，可传 song,playlist,mv,djradio,djprogram
// Msg: 内容，140 字限制，支持 emoji，@用户名
func ShareResource(data utils.RequestData, resourceID int, resourceType string, msg string) (result types.SendEventData, err error) {
	var options utils.EapiOption
	options.Path = ResourceShareAPI
	options.Url = "https://music.163.com/eapi/share/friends/resource"
	reqBodyJson := CreateResourceShareReqJson(resourceID, resourceType, msg)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
