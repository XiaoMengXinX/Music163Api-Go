package api

import (
	"encoding/json"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
	"net/http"
)

// Batch 批处理 APi
type Batch struct {
	API    map[string]interface{}
	Result string
}

// BatchAPI 被批处理的 API
type BatchAPI struct {
	Key  string
	Json string
}

// Add 添加 API
func (b *Batch) Add(apis ...BatchAPI) {
	for _, api := range apis {
		b.API[api.Key] = api.Json
	}
}

// Do 请求批处理 API
func (b *Batch) Do(data utils.RequestData) (bodyJson string, heeader http.Header, err error) {
	reqBodyJson, err := json.Marshal(b.API)
	if err != nil {
		return bodyJson, heeader, err
	}
	var options utils.EapiOption
	options.Path = "/api/batch"
	options.Url = "https://music.163.com/eapi/batch"
	options.Json = string(reqBodyJson)
	bodyJson, heeader, err = utils.EapiRequest(options, data)
	b.Result = bodyJson
	return bodyJson, heeader, err
}

// Parse 解析 Batch 的 Json 数据
func (b *Batch) Parse() map[string]string {
	jsonData := make(map[string]interface{})
	jsonMap := make(map[string]string)
	_ = json.Unmarshal([]byte(b.Result), &jsonData)
	for k, v := range jsonData {
		jsonStr, _ := json.Marshal(v)
		jsonMap[k] = string(jsonStr)
	}
	return jsonMap
}

// NewBatch 新建 Batch 对象
func NewBatch(apis ...BatchAPI) Batch {
	b := Batch{}
	b.API = make(map[string]interface{})
	b.API["e_r"] = "true"
	b.API["header"] = "{}"
	for _, api := range apis {
		b.API[api.Key] = api.Json
	}
	return b
}
