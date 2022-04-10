package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SongDetailAPI 歌曲详细信息 API
const SongDetailAPI = "/api/v3/song/detail"

// SongDetailReq SongDetail API 的 body json
type songDetailReq struct {
	C string `json:"c"`
}

type songIDs struct {
	Id int `json:"id"`
}

// CreateSongDetailReqJson 创建 获取歌单详细 请求json
func CreateSongDetailReqJson(ids []int) string {
	var songID []songIDs
	for i := 0; i < len(ids); i++ {
		songID = append(songID, songIDs{
			Id: ids[i],
		})
	}
	songIDJson, _ := json.Marshal(songID)
	reqBody := songDetailReq{
		C: string(songIDJson),
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSongDetail 获取歌曲详细信息
func GetSongDetail(data utils.RequestData, ids []int) (result types.SongsDetailData, err error) {
	var options utils.EapiOption
	options.Path = SongDetailAPI
	options.Url = "https://music.163.com/eapi/v3/song/detail"
	reqBodyJson := CreateSongDetailReqJson(ids)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
