package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

type songDetailReq struct {
	C      string `json:"c"`
	ER     string `json:"e_r"`
	Header string `json:"header"`
}

type songIDs struct {
	Id int `json:"id"`
}

// GetSongDetail 获取歌曲详细信息
func GetSongDetail(data utils.RequestData, ids []int) (result types.SongDetailData, err error) {
	var options utils.EapiOption
	options.Path = "/api/v3/song/detail"
	options.Url = "https://music.163.com/eapi/v3/song/detail"
	var songID []songIDs
	for i := 0; i < len(ids); i++ {
		songID = append(songID, songIDs{
			Id: ids[i],
		})
	}
	songIDJson, err := json.Marshal(songID)
	if err != nil {
		return types.SongDetailData{}, err
	}
	reqBody := songDetailReq{
		C:      string(songIDJson),
		Header: "{}",
		ER:     "true",
	}
	reqBodyJson, err := json.Marshal(reqBody)
	if err != nil {
		return types.SongDetailData{}, err
	}
	options.Json = string(reqBodyJson)
	resBody, err := utils.EapiRequest(options, data)
	err = json.Unmarshal([]byte(resBody), &result)
	return result, err
}
