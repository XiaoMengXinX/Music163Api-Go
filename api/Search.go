package api

import (
	"encoding/json"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

// SearchSuggestAPI 获取搜索建议 API
const SearchSuggestAPI = "/api/search/suggest/keyword"

// SearchMultiAPI 搜索多重匹配 API
const SearchMultiAPI = "/api/search/suggest/multimatch"

// SearchComplexAPI 复杂搜索 API
const SearchComplexAPI = "/api/search/complex/get/v2"

// SearchSongAPI 歌曲搜索 API
const SearchSongAPI = "/api/v1/search/song/get"

// searchSuggestReq SearchSuggest API 的 body json
type searchSuggestReq struct {
	S     string `json:"s"`
	Limit int    `json:"limit"`
}

// SearchMultiReq SearchMulti API 的 body json
type SearchMultiReq struct {
	Limit int    `json:"limit"`
	S     string `json:"s"`
}

// SearchComplexReq SearchComplex API 的 body json
type SearchComplexReq struct {
	Keyword string `json:"keyword"`
}

// SearchSongReq 歌曲搜索
type SearchSongReq struct {
	S      string `json:"s"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

// SearchSongConfig 搜索歌曲参数
type SearchSongConfig struct {
	// Keyword 搜索关键词
	Keyword string
	// Limit 最大返回结果个数 (默认为20)
	Limit int
	// Offset 偏移数量, 用于分页 (默认为0)
	Offset int
}

// CreateSearchSuggestReqJson 创建 搜索 请求json
func CreateSearchSuggestReqJson(keyword string) string {
	reqBody := searchSuggestReq{
		S:     keyword,
		Limit: 10,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// CreateSearchMultiReqJson 创建请求 body json
func CreateSearchMultiReqJson(keyword string) string {
	reqBody := SearchMultiReq{
		Limit: 5,
		S:     keyword,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// CreateSearchComplexReqJson 创建请求 body json
func CreateSearchComplexReqJson(keyword string) string {
	reqBody := SearchComplexReq{
		Keyword: keyword,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// CreateSearchSongReqJson 创建请求 body json
func CreateSearchSongReqJson(config SearchSongConfig) string {
	if config.Limit == 0 {
		config.Limit = 20
	}
	reqBody := SearchSongReq{
		S:      config.Keyword,
		Offset: config.Offset,
		Limit:  config.Limit,
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// GetSearchSuggest 获取搜索建议
func GetSearchSuggest(data utils.RequestData, keyword string) (result types.SearchSuggestData, err error) {
	var options utils.EapiOption
	options.Path = SearchSuggestAPI
	options.Url = "https://music.163.com/eapi/search/suggest/keyword"
	reqBodyJson := CreateSearchSuggestReqJson(keyword)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// SearchMultiMatch 搜索多重匹配
func SearchMultiMatch(data utils.RequestData, keyword string) (result types.SearchMultiMatchData, err error) {
	var options utils.EapiOption
	options.Path = SearchMultiAPI
	options.Url = "https://music.163.com/eapi/search/suggest/multimatch"
	reqBodyJson := CreateSearchMultiReqJson(keyword)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// SearchComplex 复杂搜索
func SearchComplex(data utils.RequestData, keyword string) (result types.SearchComplexData, err error) {
	var options utils.EapiOption
	options.Path = SearchComplexAPI
	options.Url = "https://music.163.com/eapi/search/complex/get/v2"
	reqBodyJson := CreateSearchComplexReqJson(keyword)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// SearchSong 歌曲搜索
func SearchSong(data utils.RequestData, config SearchSongConfig) (result types.SearchSongData, err error) {
	var options utils.EapiOption
	options.Path = SearchSongAPI
	options.Url = "https://music.163.com/eapi/v1/search/song/get"
	reqBodyJson := CreateSearchSongReqJson(config)
	options.Json = reqBodyJson
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
