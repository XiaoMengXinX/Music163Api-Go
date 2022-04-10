package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// ArtistHomepageAPI 获取歌手主页API
const ArtistHomepageAPI = "/api/personal/home/page/artist"

// ArtistHomepageReq 获取歌手主页API 的 body json
type ArtistHomepageReq struct {
	ArtistId int64 `json:"artistId"`
}

// CreateArtistHomepageReqJson 创建获取歌手主页API 的 body json
func CreateArtistHomepageReqJson(artistId int64) string {
	reqBody := ArtistHomepageReq{
		ArtistId: artistId,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetArtistHomepage 获取歌手主页
func GetArtistHomepage(data utils.RequestData, artistID int64) (result types.ArtistHomepageData, err error) {
	var options utils.EapiOption
	options.Path = ArtistHomepageAPI
	options.Url = "https://music.163.com/eapi/personal/home/page/artist"
	options.Json = CreateArtistHomepageReqJson(artistID)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
