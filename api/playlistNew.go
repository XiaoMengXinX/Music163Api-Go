package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// NewPlayListAPI 新建歌单 API
var NewPlayListAPI = "/api/playlist/create"

// NewPlaylistReq NewPlaylist API 的 body json
type NewPlaylistReq struct {
	Name string `json:"name"`
	// Privacy 是否设置为隐私歌单，默认否，传'10'则设置成隐私歌单
	Privacy int    `json:"privacy"`
	Header  string `json:"header"`
	ER      string `json:"e_r"`
}

// CreateNewPlaylistReqJson 创建请求 body json
func CreateNewPlaylistReqJson(name string, privacy int) string {
	reqBody := NewPlaylistReq{
		Name:    name,
		Privacy: privacy,
		Header:  "{}",
		ER:      "true",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// NewPlaylist 新建歌单
func NewPlaylist(data utils.RequestData, name string, privacy int) (result types.NewPlaylistData, err error) {
	var options utils.EapiOption
	options.Path = NewPlayListAPI
	options.Url = "https://music.163.com/eapi/playlist/create"
	reqBodyJson := CreateNewPlaylistReqJson(name, privacy)
	options.Json = reqBodyJson
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
