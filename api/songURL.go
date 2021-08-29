package api

import (
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongURLConfig 获取歌曲 URL 的参数配置
type SongURLConfig struct {
	EncodeType string
	Level      string
	Ids        []int
}

type songUrlReq struct {
	ER         string `json:"e_r"`
	EncodeType string `json:"encodeType"` // mp3, aac
	Header     string `json:"header"`
	Ids        string `json:"ids"`
	Level      string `json:"level"` // loseless, higher
}

// GetSongURL 获取歌曲 URL
func GetSongURL(data utils.RequestData, config SongURLConfig) (result types.SongURLData, err error) {
	var options utils.EapiOption
	options.Path = "/api/song/enhance/player/url/v1"
	options.Url = "https://music.163.com/eapi/song/enhance/player/url/v1"
	var IDs []string
	for i := 0; i < len(config.Ids); i++ {
		IDs = append(IDs, fmt.Sprintf("%d_0", config.Ids[i]))
	}
	IDsJson, err := json.Marshal(IDs)
	if err != nil {
		return types.SongURLData{}, err
	}
	if config.Level == "" {
		config.Level = "lossless"
	}
	if config.EncodeType == "" {
		config.EncodeType = "mp3"
	}
	reqBody := songUrlReq{
		Header:     "{}",
		ER:         "true",
		Ids:        string(IDsJson),
		EncodeType: config.EncodeType,
		Level:      config.Level,
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return types.SongURLData{}, err
	}
	options.Json = string(reqBodyJson)
	resBody, err := utils.EapiRequest(options, data)
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
