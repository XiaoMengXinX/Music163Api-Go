package types

// DelPlaylistData 删除歌单 API 的返回数据
type DelPlaylistData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Id      int64  `json:"id"`
}
