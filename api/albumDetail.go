package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// AlbumDetailAPI 获取专辑详细 API
const AlbumDetailAPI = "/api/album/v3/detail"

// AlbumDetailReq NewPlaylist API 的 body json
type AlbumDetailReq struct {
	Id       int    `json:"id"`
	CacheKey string `json:"cache_key"`
	ER       string `json:"e_r"`
	Header   string `json:"header"`
}

// CreateAlbumDetailReqJson 创建请求 body json
func CreateAlbumDetailReqJson(albumID int) string {
	reqBody := AlbumDetailReq{
		Id:       albumID,
		CacheKey: base64.StdEncoding.EncodeToString(utils.CacheKeyEncrypt(fmt.Sprintf("e_r=true&id=%d", albumID))),
		Header:   "{}",
		ER:       "true",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetAlbumDetail 获取专辑信息
func GetAlbumDetail(data utils.RequestData, albumID int) (result types.NewPlaylistData, err error) {
	var options utils.EapiOption
	options.Path = AlbumDetailAPI
	options.Url = "https://music.163.com/eapi/album/v3/detail"
	reqBodyJson := CreateAlbumDetailReqJson(albumID)
	options.Json = reqBodyJson
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
