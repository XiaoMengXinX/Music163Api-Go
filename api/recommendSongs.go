package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// RecommendSongsAPI 获取日推歌曲 API
const RecommendSongsAPI = "/api/v3/discovery/recommend/songs"

// GetRecommendSongs 获取日推歌曲
func GetRecommendSongs(data utils.RequestData) (result types.RecommendSongsData, err error) {
	var options utils.EapiOption
	options.Path = RecommendSongsAPI
	options.Url = "https://music.163.com/api/v3/discovery/recommend/songs"
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
