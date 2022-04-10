package api

import (
	"encoding/json"
	"strings"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"github.com/google/uuid"
)

// SendEventAPI 发送动态 API
const SendEventAPI = "/api/share/friends/resource"

// sendEventReq SendEvent API 的 body json
type sendEventReq struct {
	Msg        string `json:"msg"`
	Type       string `json:"type"`
	UUID       string `json:"uuid"`
	Pics       string `json:"pics"`
	AddComment string `json:"addComment"`
}

// eventPic 用于发送动态的图片数据
type eventPic struct {
	OriginId    int    `json:"originId"`
	SquareId    int    `json:"squareId"`
	RectangleId int    `json:"rectangleId"`
	Format      string `json:"format"`
}

// CreateEventReqJson 创建 发送动态 请求json
func CreateEventReqJson(text, pics string) string {
	UUID := uuid.New()
	shareConfig := sendEventReq{
		Msg:        text,
		Pics:       pics,
		Type:       "noresource",
		UUID:       strings.Replace(UUID.String(), "-", "", -1),
		AddComment: "false",
	}
	reqBodyJson, _ := json.Marshal(shareConfig)
	return string(reqBodyJson)
}

// CreateEventPicsJson 创建动态图片数据 json
func CreateEventPicsJson(picData []types.UploadEventImgData) string {
	var eventPicData []eventPic
	for i := 0; i < len(picData); i++ {
		eventPicData = append(eventPicData, eventPic{
			OriginId:    picData[i].PicInfo.OriginId,
			SquareId:    picData[i].PicInfo.SquareId,
			RectangleId: picData[i].PicInfo.RectangleId,
			Format:      picData[i].PicSubtype,
		})
	}
	resultJson, _ := json.Marshal(eventPicData)
	return string(resultJson)
}

// SendEvent 发送动态（可以带图片）
func SendEvent(data utils.RequestData, text string, picPath []string) (result types.SendEventData, err error) {
	var options utils.EapiOption
	options.Path = SendEventAPI
	options.Url = "https://music.163.com/eapi/share/friends/resource"
	if len(picPath) != 0 {
		var picData []types.UploadEventImgData
		for i := 0; i < len(picPath); i++ {
			nosToken, file, err := GetNosToken(data, picPath[i])
			if err != nil {
				return result, err
			}
			_, err = UploadFile(data, file, nosToken)
			if err != nil {
				return result, err
			}
			_, picSubtype := utils.DetectFileType(file[:32])
			uploadResult, err := UploadEventImg(data, nosToken.Result.DocId, picSubtype)
			if err != nil {
				return result, err
			}
			uploadResult.PicSubtype = picSubtype
			picData = append(picData, uploadResult)
		}
		picDataJson := CreateEventPicsJson(picData)
		options.Json = CreateEventReqJson(text, picDataJson)
	} else {
		options.Json = CreateEventReqJson(text, "[]")
	}
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
