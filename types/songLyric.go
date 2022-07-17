package types

// SongLyricData 获取歌词 API 返回数据
type SongLyricData struct {
	RawJson   string `json:"-"`
	Sgc       bool   `json:"sgc"`
	Sfy       bool   `json:"sfy"`
	Qfy       bool   `json:"qfy"`
	LyricUser struct {
		Id       int    `json:"id"`
		Status   int    `json:"status"`
		Demand   int    `json:"demand"`
		Userid   int    `json:"userid"`
		Nickname string `json:"nickname"`
		Uptime   int64  `json:"uptime"`
	} `json:"lyricUser"`
	Lrc struct {
		Version int    `json:"version"`
		Lyric   string `json:"lyric"`
	} `json:"lrc"`
	Klyric struct {
		Version int    `json:"version"`
		Lyric   string `json:"lyric"`
	} `json:"klyric"`
	Tlyric struct {
		Version int    `json:"version"`
		Lyric   string `json:"lyric"`
	} `json:"tlyric"`
	Code int `json:"code"`
}
