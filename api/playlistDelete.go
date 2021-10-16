package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// DelPlayListAPI 删除歌单 API
const DelPlayListAPI = "/api/playlist/delete"

// DelPlaylistReq DelPlaylist API 的 body json
type DelPlaylistReq struct {
	Pid    int    `json:"pid"`
	Header string `json:"header"`
	ER     string `json:"e_r"`
}

// CreateDelPlaylistReqJson 创建请求 body json
func CreateDelPlaylistReqJson(playlistID int) string {
	reqBody := DelPlaylistReq{
		Pid:    playlistID,
		Header: "{}",
		ER:     "true",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// DelPlaylist 删除歌单
func DelPlaylist(data utils.RequestData, playlistID int) (result types.DelPlaylistData, err error) {
	var options utils.EapiOption
	options.Path = DelPlayListAPI
	options.Url = "https://music.163.com/eapi/playlist/delete"
	reqBodyJson := CreateDelPlaylistReqJson(playlistID)
	options.Json = reqBodyJson
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
