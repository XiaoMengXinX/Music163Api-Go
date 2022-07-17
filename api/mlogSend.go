package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SendMlogAPI 发送 Mlog API
const SendMlogAPI = "/api/mlog/publish/v1"

type mlogSong struct {
	EndTime   int    `json:"endTime"`
	Name      string `json:"name"`
	SongId    string `json:"songId"`
	StartTime int    `json:"startTime"`
}

type picMlogJsonContent struct {
	Image     []MlogPic `json:"image"`
	NeedAudio bool      `json:"needAudio"`
	Song      mlogSong  `json:"song"`
	Text      string    `json:"text"`
}

type picMlogJson struct {
	Content picMlogJsonContent `json:"content"`
	From    int                `json:"from"`
	Type    int                `json:"type"`
}

// MlogPic Mlog 的图片对象
type MlogPic struct {
	Height int    `json:"height"`
	More   bool   `json:"more"`
	NosKey string `json:"nosKey"`
	PicKey string `json:"picKey"`
	Width  int    `json:"width"`
}

// sendMlogReqJson 发送 Mlog 请求 json
type sendMlogReqJson struct {
	Type string `json:"type"`
	Mlog string `json:"mlog"`
}

// NewMlogPicObj 创建 Mlog 图片对象
func NewMlogPicObj(picFile []byte, picData types.NosTokenData) (MlogPic, error) {
	width, height, err := utils.ImageSize(picFile)
	if err != nil {
		return MlogPic{}, err
	}
	mlogPicData := MlogPic{
		Width:  width,
		Height: height,
		More:   false,
		NosKey: fmt.Sprintf("%s/%s", picData.Data.Bucket, picData.Data.ObjectKey),
		PicKey: fmt.Sprintf("%d", picData.Data.ResourceId),
	}
	return mlogPicData, err
}

// CreatePicMlogReqJson 创建 发送mlog 请求json
func CreatePicMlogReqJson(text string, songInfo types.SongDetailData, picData []MlogPic) string {
	mlogSong := mlogSong{
		EndTime:   0,
		Name:      songInfo.Name,
		SongId:    fmt.Sprintf("%d", songInfo.Id),
		StartTime: 30000,
	}
	mlogContent := picMlogJsonContent{
		Image:     picData,
		NeedAudio: false,
		Song:      mlogSong,
		Text:      text,
	}
	mlogData := picMlogJson{
		Content: mlogContent,
		From:    0,
		Type:    1,
	}
	mlogDataJson, _ := json.Marshal(mlogData)
	mlogJson := sendMlogReqJson{
		Type: "1",
		Mlog: string(mlogDataJson),
	}
	resultJson, _ := json.Marshal(mlogJson)
	return string(resultJson)
}

// SendPicMlog 发送图片 Mlog
func SendPicMlog(data utils.RequestData, text string, songID int, picPath []string) (result types.SendMlogData, err error) {
	var options utils.EapiOption
	options.Path = SendMlogAPI
	options.Url = "https://music.163.com/eapi/mlog/publish/v1"
	var picData []MlogPic
	if len(picPath) == 0 || text == "" || songID == 0 {
		return types.SendMlogData{}, fmt.Errorf("Wrong parameters ")
	}
	songInfo, err := GetSongDetail(data, []int{songID})
	if err != nil {
		return result, err
	}
	if len(songInfo.Songs) == 0 {
		return result, fmt.Errorf("failed to get music details")
	}
	for i := 0; i < len(picPath); i++ {
		nosToken, file, err := GetMlogNosToken(data, picPath[i])
		if err != nil {
			return result, err
		}
		_, err = UploadFile(data, file, nosToken)
		if err != nil {
			return result, err
		}
		picObj, err := NewMlogPicObj(file, nosToken)
		if err != nil {
			return result, err
		}
		picData = append(picData, picObj)
	}
	options.Json = CreatePicMlogReqJson(text, songInfo.Songs[0], picData)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
