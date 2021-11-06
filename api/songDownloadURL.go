package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongDownloadURL 歌曲下载 URL API
const SongDownloadURL = "/api/song/enhance/download/url"

// SongDownloadURLReq 歌曲下载 URL 的参数配置
type SongDownloadURLReq struct {
	// Br 码率,默认设置了 999000 即最大码率
	Id int `json:"id"`
	// Ids 歌曲 ID
	Br int `json:"br"`
}

// CreateSongDownloadURLJson 创建请求 body json
func CreateSongDownloadURLJson(id int) string {
	reqBody := SongDownloadURLReq{
		Id: id,
		Br: 999000,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSongDownloadURL 获取歌曲下载 URL, 非 VIP 账号可通过此 API 获取部分歌曲的无损音质
func GetSongDownloadURL(data utils.RequestData, id int) (result types.SongDownloadURLData, err error) {
	var options utils.EapiOption
	options.Path = SongDownloadURL
	options.Url = "https://music.163.com/eapi/song/enhance/download/url"
	options.Json = CreateSongDownloadURLJson(id)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
