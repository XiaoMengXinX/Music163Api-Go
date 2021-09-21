package api

import (
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// PlaylistDetail 歌单详细信息 API
const PlaylistDetail = "/api/v6/playlist/detail"

// PlaylistDetailReq PlaylistDetailReq API 的 body json
type PlaylistDetailReq struct {
	Id     string `json:"id"`
	T      string `json:"t"`
	N      string `json:"n"`
	S      string `json:"s"`
	Header string `json:"header"`
	ER     string `json:"e_r"`
}

// CreatePlaylistDetailReqJson 创建请求 body json
func CreatePlaylistDetailReqJson(id int) string {
	reqBody := PlaylistDetailReq{
		Id:     fmt.Sprintf("%d", id),
		T:      "0",
		N:      "50",
		S:      "5",
		Header: "{}",
		ER:     "true",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetPlaylistDetail 获取歌单详细信息
func GetPlaylistDetail(data utils.RequestData, id int) (result types.PlaylistDetailData, err error) {
	var options utils.EapiOption
	options.Path = PlaylistDetail
	options.Url = "https://music.163.com/eapi/v6/playlist/detail"
	reqBodyJson := CreatePlaylistDetailReqJson(id)
	options.Json = reqBodyJson
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
