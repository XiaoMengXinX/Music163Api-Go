package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UserPlaylistAPI 获取用户歌单 API
const UserPlaylistAPI = "/api/user/playlist"

// UserPlaylistConfig 获取用户歌单 API 的参数配置
type UserPlaylistConfig struct {
	// UserID 用户 ID
	UserID int
	// Offset 偏移数量，用于分页 , 如 :( 页数 -1)*30, 其中 30 为 limit 的值 , 默认为 0
	Offset int
	// Limit , 默认为 30
	Limit int
}

// userPlaylistReq UserPlaylist API 的 body json
type userPlaylistReq struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Uid    int `json:"uid"`
}

// CreateUserPlaylistReqJson 创建 获取用户歌单 请求json
func CreateUserPlaylistReqJson(config UserPlaylistConfig) string {
	if config.Limit == 0 {
		config.Limit = 30
	}
	reqBody := userPlaylistReq{
		Offset: config.Offset,
		Limit:  config.Limit,
		Uid:    config.UserID,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetUserPlaylist 获取用户歌单
func GetUserPlaylist(data utils.RequestData, config UserPlaylistConfig) (result types.UserPlaylistData, err error) {
	var options utils.EapiOption
	options.Path = UserPlaylistAPI
	options.Url = "https://music.163.com/eapi/user/playlist"
	reqBodyJson := CreateUserPlaylistReqJson(config)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
