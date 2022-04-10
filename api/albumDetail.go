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

// albumDetailReq NewPlaylist API 的 body json
type albumDetailReq struct {
	Id       int    `json:"id"`
	CacheKey string `json:"cache_key"`
}

// CreateAlbumDetailReqJson 创建 获取专辑详细 请求json
func CreateAlbumDetailReqJson(albumID int) string {
	reqBody := albumDetailReq{
		Id:       albumID,
		CacheKey: base64.StdEncoding.EncodeToString(utils.CacheKeyEncrypt(fmt.Sprintf("id=%d", albumID))),
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
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
