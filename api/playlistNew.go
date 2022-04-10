package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// NewPlayListAPI 新建歌单 API
const NewPlayListAPI = "/api/playlist/create"

// NewPlaylistReq NewPlaylist API 的 body json
type NewPlaylistReq struct {
	Name    string `json:"name"`
	Privacy int    `json:"privacy"`
}

// CreateNewPlaylistReqJson 创建 新建歌单 请求json
func CreateNewPlaylistReqJson(name string, isPrivate bool) string {
	reqBody := NewPlaylistReq{Name: name, Privacy: 0}
	if isPrivate {
		reqBody.Privacy = 10
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// NewPlaylist 新建歌单
func NewPlaylist(data utils.RequestData, name string, isPrivate bool) (result types.NewPlaylistData, err error) {
	var options utils.EapiOption
	options.Path = NewPlayListAPI
	options.Url = "https://music.163.com/eapi/playlist/create"
	reqBodyJson := CreateNewPlaylistReqJson(name, isPrivate)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
