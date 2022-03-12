package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// ArtistDetailAPI 获取歌手详情API
const ArtistDetailAPI = "/api/personal/home/page/artist"

// ArtistDetailReq 获取歌手详情API 的 body json
type ArtistDetailReq struct {
	ArtistId int64 `json:"artistId"`
}

// CreateArtistDetailReqJson 创建获取歌手详情API 的 body json
func CreateArtistDetailReqJson(artistId int64) string {
	reqBody := ArtistDetailReq{
		ArtistId: artistId,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetArtistDetail 获取歌手详情
func GetArtistDetail(data utils.RequestData, artistID int64) (result types.ArtistDetailData, err error) {
	var options utils.EapiOption
	options.Path = ArtistDetailAPI
	options.Url = "https://music.163.com/eapi/personal/home/page/artist"
	options.Json = CreateArtistDetailReqJson(artistID)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
