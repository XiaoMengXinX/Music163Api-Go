package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongShareAPI 分享歌曲 API
const SongShareAPI = "/api/music/songshare/share/property"

// songShareReq SongShare API 的 body json
type songShareReq struct {
	SongId int `json:"songId"`
}

// CreateSongShareReqJson 创建 分享歌曲 请求json
func CreateSongShareReqJson(musicID int) string {
	reqBody := songShareReq{
		SongId: musicID,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// SongShare 分享歌曲，用于音乐人任务
func SongShare(data utils.RequestData, musicID int) (result types.SongShareData, err error) {
	var options utils.EapiOption
	options.Path = SongShareAPI
	options.Url = "https://music.163.com/eapi/song/share"
	reqBodyJson := CreateSongShareReqJson(musicID)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
