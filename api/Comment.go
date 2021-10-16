package api

import (
	"encoding/json"
	"fmt"
	"github.com/XiaoMengXinX/Music163Api-Go/types"
	"github.com/XiaoMengXinX/Music163Api-Go/utils"
)

const (
	// ReplyCommentAPI 回复评论 API
	ReplyCommentAPI = "/api/v1/resource/comments/reply"
	// DelCommentAPI 删除评论 API
	DelCommentAPI = "/api/resource/comments/delete"
	// AddCommentAPI 新增评论 API
	AddCommentAPI = "/api/v1/resource/comments/add"
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

// CommentConfig 回复/发送/删除评论参数.
// ResType 0: 歌曲(默认), 1: mv, 2: 歌单, 3: 专辑, 4: 电台, 5: 视频, 6: Mlog, 7: 动态.
// ForwardEvent 只在发送或回复评论时有效.
type CommentConfig struct {
	// ResType 0: 歌曲(默认), 1: mv, 2: 歌单, 3: 专辑, 4: 电台, 5: 视频, 6: Mlog, 7: 动态
	ResType int
	// ResID 资源 ID
	ResID int
	// CommentID 评论 ID
	CommentID int
	// ForwardEvent 是否转发到动态, 只在发送或回复评论时有效
	ForwardEvent bool
	// Content 评论内容
	Content string
}

// CommentReq 评论 API 的 body json
type CommentReq struct {
	ForwardEvent string `json:"forwardEvent"`
	ThreadId     string `json:"threadId"`
	CommentId    string `json:"commentId"`
	Content      string `json:"content"`
	Header       string `json:"header"`
	ER           string `json:"e_r"`
}

// CreateCommentJson 创建请求 body json
func CreateCommentJson(config CommentConfig) string {
	resExt := parseRestype(config.ResType)
	reqBody := CommentReq{
		Content:   config.Content,
		ThreadId:  fmt.Sprintf("%s_%d", resExt, config.ResID),
		CommentId: fmt.Sprintf("%d", config.CommentID),
		Header:    "{}",
		ER:        "true",
	}
	if config.ForwardEvent {
		reqBody.ForwardEvent = "1"
	} else {
		reqBody.ForwardEvent = "0"
	}
	reqBodyJson, _ := json.Marshal(reqBody)
	return string(reqBodyJson)
}

func parseRestype(resType int) string {
	var resExt string
	switch resType {
	case 0:
		resExt = ResTypeMusic
	case 1:
		resExt = ResTypeMV
	case 2:
		resExt = ResTypePlaylist
	case 3:
		resExt = ResTypeAlbum
	case 4:
		resExt = ResTypeRadio
	case 5:
		resExt = ResTypeVideo
	case 6:
		resExt = ResTypeMlog
	case 7:
		resExt = ResTypeEvent
	}
	return resExt
}

// AddComment 新增评论
func AddComment(data utils.RequestData, config CommentConfig) (result types.AddCommentData, err error) {
	var options utils.EapiOption
	options.Path = AddCommentAPI
	options.Url = "https://music.163.com/eapi/v1/resource/comments/add"
	options.Json = CreateCommentJson(config)
	resBody, _, err := utils.EapiRequest(options, data)
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
	resBody, _, err := utils.EapiRequest(options, data)
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
	resBody, _, err := utils.EapiRequest(options, data)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(resBody), &result)
	result.RawJson = resBody
	return result, err
}
