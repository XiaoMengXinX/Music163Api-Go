package api

import (
	"encoding/json"
	"net/http"

	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// Batch 批处理 APi
type Batch struct {
	API    map[string]interface{}
	Result string
	Header http.Header
	Error  error
}

// BatchAPI 被批处理的 API
type BatchAPI struct {
	Key  string
	Json string
}

// Add 添加 API
func (b *Batch) Add(apis ...BatchAPI) *Batch {
	for _, api := range apis {
		b.API[api.Key] = api.Json
	}
	return b
}

// Do 请求批处理 API
func (b *Batch) Do(data utils.RequestData) *Batch {
	reqBodyJson, err := json.Marshal(b.API)
	if err != nil {
		b.Error = err
		return b
	}
	var options utils.EapiOption
	options.Path = "/api/batch"
	options.Url = "https://music.163.com/eapi/batch"
	options.Json = string(reqBodyJson)
	b.Result, b.Header, b.Error = utils.ApiRequest(options, data)
	return b
}

// Parse 解析 Batch 的 Json 数据
func (b *Batch) Parse() (*Batch, map[string]string) {
	jsonData := make(map[string]interface{})
	jsonMap := make(map[string]string)
	_ = json.Unmarshal([]byte(b.Result), &jsonData)
	for k, v := range jsonData {
		jsonStr, _ := json.Marshal(v)
		jsonMap[k] = string(jsonStr)
	}
	return b, jsonMap
}

// NewBatch 新建 Batch 对象
func NewBatch(apis ...BatchAPI) *Batch {
	b := &Batch{}
	b.API = make(map[string]interface{})
	for _, api := range apis {
		b.API[api.Key] = api.Json
	}
	return b
}
