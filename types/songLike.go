package types

// SongLikeData 红心/取消红心歌曲返回数据
type SongLikeData struct {
	RawJson    string
	PlaylistId int `json:"playlistId"`
	Code       int `json:"code"`
}
