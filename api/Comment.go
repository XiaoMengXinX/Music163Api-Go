package api

import (
	"encoding/json"
	"fmt"

	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

const (
	// ResTypeMusic 歌曲
	ResTypeMusic = "R_SO_4"
	// ResTypeMV MV
	ResTypeMV = "R_MV_5"
	// ResTypePlaylist 歌单
	ResTypePlaylist = "A_PL_0"
	// ResTypeAlbum 专辑
	ResTypeAlbum = "R_AL_3"
	// ResTypeRadio 电台
	ResTypeRadio = "A_DJ_1"
	// ResTypeVideo 视频
	ResTypeVideo = "R_VI_62"
	// ResTypeMlog Mlog
	ResTypeMlog = "R_MLOG_1001"
	// ResTypeEvent 动态
	ResTypeEvent = "A_EV_2"
)

// ReplyCommentAPI 回复评论 API
const ReplyCommentAPI = "/api/v1/resource/comments/reply"

// DelCommentAPI 删除评论 API
const DelCommentAPI = "/api/resource/comments/delete"

// AddCommentAPI 新增评论 API
const AddCommentAPI = "/api/v1/resource/comments/add"

// GetCommentAPI 获取评论 API
const GetCommentAPI = "/api/v2/resource/comments"

// CommentConfig 回复/发送/删除评论参数.
type CommentConfig struct {
	// ResType 对应常量 ResTypeXXX
	ResType string
	// ResID 资源 ID
	ResID int
	// CommentID 评论 ID
	CommentID int
	// ForwardEvent 是否转发到动态, 只在发送或回复评论时有效
	ForwardEvent bool
	// Content 评论内容
	Content string
}

// commentReq 评论 API 的 body json
type commentReq struct {
	ForwardEvent string `json:"forwardEvent"`
	ThreadId     string `json:"threadId"`
	CommentId    string `json:"commentId"`
	Content      string `json:"content"`
}

// GetCommentConfig 获取评论参数
type GetCommentConfig struct {
	// ResType 对应常量 ResTypeXXX
	ResType string
	// ResID 资源 ID
	ResID int
	// PageNo 分页参数,第N页,默认为1
	PageNo int
	// PageSize 分页参数,每页多少条数据,默认20
	PageSize int
	// SortType 排序方式,99:按推荐排序,2:按热度排序,3:按时间排序,默认为99
	SortType int
	// Cursor 当sortType为3时且页数不是第一页时需传入,值为上一条数据的time
	Cursor int
}

// getCommentReq 获取评论 API 的 body json
type getCommentReq struct {
	Cursor    int    `json:"cursor"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
	ShowInner bool   `json:"showInner"`
	SortType  int    `json:"sortType"`
	ThreadId  string `json:"threadId"`
}

// CreateCommentJson 创建 发送/回复/删除评论 请求json
func CreateCommentJson(config CommentConfig) string {
	reqBody := commentReq{
		Content:   config.Content,
		ThreadId:  fmt.Sprintf("%s_%d", config.ResType, config.ResID),
		CommentId: fmt.Sprintf("%d", config.CommentID),
	}
	if config.ForwardEvent {
		reqBody.ForwardEvent = "1"
	} else {
		reqBody.ForwardEvent = "0"
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// CreateGetCommentJson 创建 获取评论 请求json
func CreateGetCommentJson(config GetCommentConfig) string {
	reqBody := getCommentReq{
		ShowInner: false,
		ThreadId:  fmt.Sprintf("%s_%d", config.ResType, config.ResID),
	}
	if config.PageNo == 0 {
		reqBody.PageNo = 1
	} else {
		reqBody.PageNo = config.PageNo
	}
	if config.PageSize == 0 {
		reqBody.PageSize = 20
	} else {
		reqBody.PageSize = config.PageSize
	}
	if config.SortType == 0 {
		reqBody.SortType = 99
	} else {
		reqBody.SortType = config.SortType
	}
	if config.SortType == 3 && config.PageNo != 1 {
		reqBody.Cursor = config.Cursor
	} else {
		reqBody.Cursor = 0
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

// AddComment 新增评论
func AddComment(data utils.RequestData, config CommentConfig) (result types.AddCommentData, err error) {
	var options utils.EapiOption
	options.Path = AddCommentAPI
	options.Url = "https://music.163.com/eapi/v1/resource/comments/add"
	options.Json = CreateCommentJson(config)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// ReplyComment 回复评论
func ReplyComment(data utils.RequestData, config CommentConfig) (result types.ReplyCommentData, err error) {
	var options utils.EapiOption
	options.Path = ReplyCommentAPI
	options.Url = "https://music.163.com/eapi/v1/resource/comments/reply"
	options.Json = CreateCommentJson(config)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// DelComment 删除评论
func DelComment(data utils.RequestData, config CommentConfig) (result types.DelCommentData, err error) {
	var options utils.EapiOption
	options.Path = DelCommentAPI
	options.Url = "https://music.163.com/eapi/resource/comments/delete"
	options.Json = CreateCommentJson(config)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}

// GetComment 获取评论
func GetComment(data utils.RequestData, config GetCommentConfig) (result types.GetCommentData, err error) {
	var options utils.EapiOption
	options.Path = GetCommentAPI
	options.Url = "https://music.163.com/eapi/v2/resource/comments"
	options.Json = CreateGetCommentJson(config)
	resBody, _, err := utils.ApiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
