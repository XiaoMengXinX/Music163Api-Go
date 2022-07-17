package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// UploadNode 上传文件节点
var UploadNode string

// GetUploadNode 获取上传加速节点地址
func GetUploadNode() (result types.UploadNodeData, err error) {
	url := "https://wanproxy.127.net/lbs?version=1.0&bucketname=cloudmusic"
	resBody, err := utils.RawRequest(url, utils.RequestData{})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// UploadFile 上传文件
func UploadFile(data utils.RequestData, file []byte, nosToken types.NosTokenData) (result types.UploadFileData, err error) {
	if UploadNode == "" {
		nodeData, err := GetUploadNode()
		if err != nil {
			return result, err
		}
		if len(nodeData.Upload) != 0 {
			UploadNode = nodeData.Upload[0]
		} else {
			return result, fmt.Errorf("Failed to get upload acceleration node ")
		}
	}
	fileType, fileSubtype := utils.DetectFileType(file[:32])
	var url string
	switch nosToken.Type {
	case 0:
		url = fmt.Sprintf("%s/%s/%s?version=1.0&offset=0&complete=true", UploadNode, nosToken.Result.Bucket, nosToken.Result.ObjectKey)
		data.Headers = utils.Headers{
			{
				Name:  "x-nos-token",
				Value: nosToken.Result.Token,
			}, {
				Name:  "Content-Type",
				Value: fmt.Sprintf("%s/%s", fileType, fileSubtype),
			},
		}
	case 1:
		url = fmt.Sprintf("%s/%s/%s?version=1.0&offset=0&complete=true", UploadNode, nosToken.Data.Bucket, nosToken.Data.ObjectKey)
		data.Headers = utils.Headers{
			{
				Name:  "x-nos-token",
				Value: nosToken.Data.Token,
			}, {
				Name:  "Content-Type",
				Value: fmt.Sprintf("%s/%s", fileType, fileSubtype),
			},
		}
	}
	data.Body = string(file)
	resBody, err := utils.RawRequest(url, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// UploadEventImgAPI 获取用于发送动态的图片信息
const UploadEventImgAPI = "/api/upload/event/img/v1"

// EventImgReq UploadEventImgAPI 的请求 body json
type EventImgReq struct {
	ImgID  string `json:"imgid"`
	Format string `json:"format"`
}

// CreateEventImgReqJson 创建 动态上传图片 请求json
func CreateEventImgReqJson(imgID, imgSubtype string) string {
	reqBody := EventImgReq{
		ImgID:  imgID,
		Format: imgSubtype,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// UploadEventImg 获取用于发送动态的图片信息
func UploadEventImg(data utils.RequestData, imgID, imgType string) (result types.UploadEventImgData, err error) {
	var options utils.EapiOption
	options.Path = UploadEventImgAPI
	options.Url = "https://music.163.com/eapi/upload/event/img/v1"
	options.Json = CreateEventImgReqJson(imgID, imgType)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
