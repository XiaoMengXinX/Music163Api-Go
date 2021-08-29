package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

const (
	// UserSetting 获取用户设置 API
	UserSetting = "/api/user/setting"
)

// Batch 批处理 APi
type Batch struct {
	API map[string]interface{}
}

// BatchAPI 被批处理的 API
type BatchAPI struct {
	Key  string
	Json string
}

// Init 初始化批处理
func (b *Batch) Init() {
	b.API = make(map[string]interface{})
	b.API["e_r"] = "true"
	b.API["header"] = "{}"
}

// Add 添加 API
func (b *Batch) Add(api BatchAPI) {
	b.API[api.Key] = api.Json
}

// Do 请求批处理 API
func (b *Batch) Do(data utils.RequestData) (string, error) {
	reqBodyJson, err := json.Marshal(b.API)
	if err != nil {
		return "", err
	}
	var options utils.EapiOption
	options.Path = "/api/batch"
	options.Url = "https://music.163.com/eapi/batch"
	options.Json = string(reqBodyJson)
	resBody, err := utils.EapiRequest(options, data)
	return resBody, err
}
