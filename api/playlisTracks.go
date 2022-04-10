package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// PlaylistTracksAPI 管理歌单歌曲 API
const PlaylistTracksAPI = "/api/playlist/manipulate/tracks"

// playListTracksReq PlaylistTracks API 的 body json
type playListTracksReq struct {
	TrackIds []int  `json:"trackIds"`
	Pid      int    `json:"pid"`
	Op       string `json:"op"`
	Imme     string `json:"imme"`
}

// CreatePlayListTracksReqJson 创建 歌单操作 请求json, operation: 1为添加, 0或其他值为删除
func CreatePlayListTracksReqJson(musicIds []int, playlistID, operation int) string {
	reqBody := playListTracksReq{
		TrackIds: musicIds,
		Pid:      playlistID,
		Imme:     "true",
	}
	switch operation {
	case 1:
		reqBody.Op = "add"
	default:
		reqBody.Op = "del"
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// AddPlaylistTracks 向歌单中添加歌曲
func AddPlaylistTracks(data utils.RequestData, musicIds []int, playlistID int) (result types.PlaylistTracksData, err error) {
	var options utils.EapiOption
	options.Path = PlaylistTracksAPI
	options.Url = "https://music.163.com/eapi/playlist/manipulate/tracks"
	reqBodyJson := CreatePlayListTracksReqJson(musicIds, playlistID, 1)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// DelPlaylistTracks 删除歌单歌曲
func DelPlaylistTracks(data utils.RequestData, trackIds []int, playlistID int) (result types.PlaylistTracksData, err error) {
	var options utils.EapiOption
	options.Path = PlaylistTracksAPI
	options.Url = "https://music.163.com/eapi/playlist/manipulate/tracks"
	reqBodyJson := CreatePlayListTracksReqJson(trackIds, playlistID, 0)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
