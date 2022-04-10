package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongUrlAPI 歌曲 URl API
const SongUrlAPI = "/api/song/enhance/player/url/v1"

// SongURLConfig 获取歌曲 URL 的参数配置
type SongURLConfig struct {
	// EncodeType 编码类型, 可选 "mp3", "aac"
	EncodeType string
	// Level 音质等级, 可选 "lossless", "higher", "standard"
	Level string
	// Ids 歌曲 ID
	Ids []int
}

// songUrlReq SongURL API 的 body json
type songUrlReq struct {
	EncodeType string `json:"encodeType"`
	Ids        string `json:"ids"`
	Level      string `json:"level"`
}

// CreateSongURLJson 创建 获取歌曲试听URL 请求json
func CreateSongURLJson(config SongURLConfig) string {
	var IDs []string
	for i := 0; i < len(config.Ids); i++ {
		IDs = append(IDs, fmt.Sprintf("%d", config.Ids[i]))
	}
	IDsJson, _ := json.Marshal(IDs)
	if config.Level == "" {
		config.Level = "lossless"
	}
	if config.EncodeType == "" {
		config.EncodeType = "mp3"
	}
	reqBody := songUrlReq{
		Ids:        string(IDsJson),
		EncodeType: config.EncodeType,
		Level:      config.Level,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSongURL 获取歌曲试听 URL
func GetSongURL(data utils.RequestData, config SongURLConfig) (result types.SongsURLData, err error) {
	var options utils.EapiOption
	options.Path = SongUrlAPI
	options.Url = "https://music.163.com/eapi/song/enhance/player/url/v1"
	options.Json = CreateSongURLJson(config)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
