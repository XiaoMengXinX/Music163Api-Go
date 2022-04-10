package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// ProgramDetailAPI 节目详情API
const ProgramDetailAPI = `/api/dj/program/detail`

type programDetailReq struct {
	ID string `json:"id"`
}

// CreateProgramDetailReqJson 创建获取节目详情请求json
func CreateProgramDetailReqJson(id int) string {
	reqBody := programDetailReq{
		ID: fmt.Sprintf("%d", id),
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetProgramDetail 获取节目详情
func GetProgramDetail(data utils.RequestData, id int) (result types.ProgramDetailData, err error) {
	var options utils.EapiOption
	options.Path = ProgramDetailAPI
	options.Url = "https://music.163.com/eapi/dj/program/detail"
	reqBodyJson := CreateProgramDetailReqJson(id)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
