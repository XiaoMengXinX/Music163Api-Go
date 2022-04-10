package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// ShortUrlAPI 短链接 API
const ShortUrlAPI = "/api/middle/shorturl/generate"

// ShortUrlReq ShortUrl API 的 body json
type shortUrlReq struct {
	Url string `json:"url"`
}

// CreateShortUrlJson 创建 获取短链接 请求json
func CreateShortUrlJson(url string) string {
	reqBody := shortUrlReq{
		Url: url,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// NewShortURL 获取短连接 (只支持 *.163.com *.126.net)
func NewShortURL(data utils.RequestData, url string) (result types.ShortURLData, err error) {
	var options utils.EapiOption
	options.Path = ShortUrlAPI
	options.Url = "https://music.163.com/eapi/middle/shorturl/generate"
	options.Json = CreateShortUrlJson(url)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
