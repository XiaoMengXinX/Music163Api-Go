package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongDetail 歌曲详细信息 API
const SongDetail = "/api/v3/song/detail"

// SongDetailReq SongDetail API 的 body json
type SongDetailReq struct {
	C      string `json:"c"`
	ER     string `json:"e_r"`
	Header string `json:"header"`
}

type songIDs struct {
	Id int `json:"id"`
}

// CreateSongDetailReqJson 创建请求 body json
func CreateSongDetailReqJson(ids []int) string {
	var songID []songIDs
	for i := 0; i < len(ids); i++ {
		songID = append(songID, songIDs{
			Id: ids[i],
		})
	}
	songIDJson, _ := json.Marshal(songID)
	reqBody := SongDetailReq{
		C:      string(songIDJson),
		Header: "{}",
		ER:     "true",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSongDetail 获取歌曲详细信息
func GetSongDetail(data utils.RequestData, ids []int) (result types.SongsDetailData, err error) {
	var options utils.EapiOption
	options.Path = SongDetail
	options.Url = "https://music.163.com/eapi/v3/song/detail"
	reqBodyJson := CreateSongDetailReqJson(ids)
	options.Json = reqBodyJson
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
