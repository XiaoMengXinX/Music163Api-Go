package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongLyricAPI 获取歌词 API
const SongLyricAPI = "/api/song/lyric"

// songLyricReq SongLyric API 的 body json
type songLyricReq struct {
	Id int `json:"id"`
	Lv int `json:"lv"`
	Kv int `json:"kv"`
	Tv int `json:"tv"`
}

// CreateSongLyricReqJson 创建 获取歌曲歌词 请求json
func CreateSongLyricReqJson(id int) string {
	reqBody := songLyricReq{
		Id: id,
		Lv: -1,
		Kv: -1,
		Tv: -1,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSongLyric 获取歌曲歌词
func GetSongLyric(data utils.RequestData, id int) (result types.SongLyricData, err error) {
	var options utils.EapiOption
	options.Path = SongLyricAPI
	options.Url = "https://music.163.com/eapi/song/lyric"
	reqBodyJson := CreateSongLyricReqJson(id)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
