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

// SearchSongGetAPI 旧版歌曲搜索 API
const SearchSongGetAPI = "/api/search/get"

// CloudSearchAPI 云搜索 API
const CloudSearchAPI = "/api/cloudsearch/pc"

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

// SearchSongGetReq 带类型的歌曲搜索请求
type SearchSongGetReq struct {
	S      string `json:"s"`
	Type   int    `json:"type"`
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
	// Type 搜索类型，1 为单曲 (默认为1，仅 SearchSongGet / CloudSearchSong 使用)
	Type int
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

// CreateSearchSongGetReqJson 创建带类型的歌曲搜索请求 body json
func CreateSearchSongGetReqJson(config SearchSongConfig) string {
	if config.Limit == 0 {
		config.Limit = 20
	}
	if config.Type == 0 {
		config.Type = 1
	}
	reqBody := SearchSongGetReq{
		S:      config.Keyword,
		Type:   config.Type,
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

func doSearch(data utils.RequestData, apiPath, url, reqBody string, dst interface{}) (string, error) {
	resBody, _, err := utils.ApiRequest(utils.EapiOption{
		Path: apiPath,
		Url:  url,
		Json: reqBody,
	}, data)
	if err != nil {
		return "", err
	}
	return resBody, json.Unmarshal([]byte(resBody), dst)
}

// SearchSong 歌曲搜索
func SearchSong(data utils.RequestData, config SearchSongConfig) (result types.SearchSongData, err error) {
	resBody, err := doSearch(data, SearchSongAPI, "https://music.163.com/eapi/v1/search/song/get", CreateSearchSongReqJson(config), &result)
	result.RawJson = resBody
	return result, err
}

// SearchSongGet 歌曲搜索 (/api/search/get)
func SearchSongGet(data utils.RequestData, config SearchSongConfig) (result types.SearchSongData, err error) {
	resBody, err := doSearch(data, SearchSongGetAPI, "https://music.163.com/eapi/search/get", CreateSearchSongGetReqJson(config), &result)
	result.RawJson = resBody
	return result, err
}

// CloudSearchSong 云搜索歌曲 (/api/cloudsearch/pc)
func CloudSearchSong(data utils.RequestData, config SearchSongConfig) (result types.CloudSearchSongData, err error) {
	resBody, err := doSearch(data, CloudSearchAPI, "https://music.163.com/eapi/cloudsearch/pc", CreateSearchSongGetReqJson(config), &result)
	result.RawJson = resBody
	return result, err
}
