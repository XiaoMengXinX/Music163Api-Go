package types

// PlaylistTracksData 歌单添加/删除歌曲返回数据
type PlaylistTracksData struct {
	RawJson    string        `json:"-"`
	OfflineIds []interface{} `json:"offlineIds"`
	TrackIds   string        `json:"trackIds"`
	Code       int           `json:"code"`
	Count      int           `json:"count"`
	CloudCount int           `json:"cloudCount"`
}
