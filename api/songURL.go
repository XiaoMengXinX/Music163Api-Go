package api

import (
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongURL 歌曲 URl API
const SongURL = "/api/song/enhance/player/url/v1"

// SongURLConfig 获取歌曲 URL 的参数配置
type SongURLConfig struct {
	EncodeType string
	Level      string
	Ids        []int
}

// SongUrlReq SongURL API 的 body json
type SongUrlReq struct {
	ER         string `json:"e_r"`
	EncodeType string `json:"encodeType"` // mp3, aac
	Header     string `json:"header"`
	Ids        string `json:"ids"`
	Level      string `json:"level"` // loseless, higher
}

// CreateSongURLJSON 创建请求 body json
func CreateSongURLJSON(config SongURLConfig) string {
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
	reqBody := SongUrlReq{
		Header:     "{}",
		ER:         "true",
		Ids:        string(IDsJson),
		EncodeType: config.EncodeType,
		Level:      config.Level,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSongURL 获取歌曲 URL
func GetSongURL(data utils.RequestData, config SongURLConfig) (result types.SongURLData, err error) {
	var options utils.EapiOption
	options.Path = SongURL
	options.Url = "https://music.163.com/eapi/song/enhance/player/url/v1"
	reqBodyJson := CreateSongURLJSON(config)
	options.Json = reqBodyJson
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
