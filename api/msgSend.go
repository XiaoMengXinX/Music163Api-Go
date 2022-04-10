package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SendMsgAPI 发送私信 API
const SendMsgAPI = "/api/msg/private/send"

// sendMsgReq 发送私信 API 的 body json
type sendMsgReq struct {
	UserIds string `json:"userIds"`
	Msg     string `json:"msg"`
	Type    string `json:"type"`
}

// msgPicInfo 发送私信的图片信息
type msgPicInfo struct {
	PicIdStr string `json:"picIdStr"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Format   string `json:"format"`
	Type     int    `json:"type"`
}

// sendPicMsgReq 发送图片私信 API 的 body json
type sendPicMsgReq struct {
	UserIds string `json:"userIds"`
	PicInfo string `json:"picinfo"`
	Type    string `json:"type"`
}

// CreateTextMsgReqJson 创建 发送私信 请求json
func CreateTextMsgReqJson(userIDs []int, msg string) string {
	var userID []int
	for i := 0; i < len(userIDs); i++ {
		userID = append(userID, userIDs[i])
	}
	userIDJson, _ := json.Marshal(userID)
	reqBody := sendMsgReq{
		UserIds: string(userIDJson),
		Msg:     msg,
		Type:    "text",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// CreatePicMsgReqJson 创建私信图片json数据
func CreatePicMsgReqJson(userIDs []int, picData types.NosTokenData, picFile []byte) (string, error) {
	var userID []int
	for i := 0; i < len(userIDs); i++ {
		userID = append(userID, userIDs[i])
	}
	userIDJson, _ := json.Marshal(userID)
	width, height, err := utils.ImageSize(picFile)
	if err != nil {
		return "", err
	}
	_, format := utils.DetectFileType(picFile)
	picInfo := msgPicInfo{
		PicIdStr: picData.Result.DocId,
		Width:    width,
		Height:   height,
		Format:   format,
		Type:     0,
	}
	picInfoJson, _ := json.Marshal(picInfo)
	reqBody := sendPicMsgReq{
		UserIds: string(userIDJson),
		PicInfo: string(picInfoJson),
		Type:    "pic",
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson), nil
}

// SendTextMsg 发送文本私信
func SendTextMsg(data utils.RequestData, userIDs []int, text string) (result types.SendMsgData, err error) {
	var options utils.EapiOption
	options.Path = SendMsgAPI
	options.Url = "https://music.163.com/eapi/msg/private/send"
	options.Json = CreateTextMsgReqJson(userIDs, text)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// SendPicMsg 发送图片私信
func SendPicMsg(data utils.RequestData, userIDs []int, picPath string) (result types.SendMsgData, err error) {
	var options utils.EapiOption
	options.Path = SendMsgAPI
	options.Url = "https://music.163.com/eapi/msg/private/send"
	nosToken, file, err := GetNosToken(data, picPath)
	if err != nil {
		return result, err
	}
	_, err = UploadFile(data, file, nosToken)
	if err != nil {
		return result, err
	}
	options.Json, err = CreatePicMsgReqJson(userIDs, nosToken, file)
	if err != nil {
		return result, err
	}
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
