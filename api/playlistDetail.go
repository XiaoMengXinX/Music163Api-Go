package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// PlaylistDetailAPI 歌单详细信息 API
const PlaylistDetailAPI = "/api/v6/playlist/detail"

// PlaylistDetailReq PlaylistDetail API 的 body json
type playlistDetailReq struct {
	ID string `json:"id"`
	T  string `json:"t"`
	N  string `json:"n"`
	S  string `json:"s"`
}

// CreatePlaylistDetailReqJson 创建 获取歌单 请求json
func CreatePlaylistDetailReqJson(id int) string {
	reqBody := playlistDetailReq{
		ID: fmt.Sprintf("%d", id),
		T:  "0",
		N:  "50",
		S:  "5",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetPlaylistDetail 获取歌单详细信息
func GetPlaylistDetail(data utils.RequestData, id int) (result types.PlaylistDetailData, err error) {
	var options utils.EapiOption
	options.Path = PlaylistDetailAPI
	options.Url = "https://music.163.com/eapi/v6/playlist/detail"
	reqBodyJson := CreatePlaylistDetailReqJson(id)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
