package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
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
func (b *Batch) Add(apis ...BatchAPI) {
	for _, api := range apis {
		b.API[api.Key] = api.Json
	}
}

// Do 请求批处理 API
func (b *Batch) Do(data utils.RequestData) (resBody, header string, err error) {
	reqBodyJson, err := json.Marshal(b.API)
	if err != nil {
		return resBody, header, err
	}
	var options utils.EapiOption
	options.Path = "/api/batch"
	options.Url = "https://music.163.com/eapi/batch"
	options.Json = string(reqBodyJson)
	resBody, header, err = utils.EapiRequest(options, data)
	return resBody, header, err
}

// NewBatch 新建 Batch 对象
func NewBatch(apis ...BatchAPI) Batch {
	b := Batch{}
	b.Init()
	for _, api := range apis {
		b.API[api.Key] = api.Json
	}
	return b
}
