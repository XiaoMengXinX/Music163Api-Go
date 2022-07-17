package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongLikeAPI 添加/删除红心歌曲 API
const SongLikeAPI = "/api/song/like"

// songLikeReq SongLike API 的 body json
type songLikeReq struct {
	TrackId int  `json:"trackId"`
	Like    bool `json:"like"`
}

// CreateSongLikeReqJson 创建 红心歌曲 请求json
func CreateSongLikeReqJson(musicID int, isLike bool) string {
	reqBody := songLikeReq{
		TrackId: musicID,
		Like:    isLike,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// SongLike 红心/取消红心歌曲
func SongLike(data utils.RequestData, musicID int, isLike bool) (result types.SongLikeData, err error) {
	var options utils.EapiOption
	options.Path = SongLikeAPI
	options.Url = "https://music.163.com/eapi/song/like"
	reqBodyJson := CreateSongLikeReqJson(musicID, isLike)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
