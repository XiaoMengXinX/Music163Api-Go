package types

// RegisterAnonData 匿名登录 API 的返回数据
type RegisterAnonData struct {
	RawJson    string `json:"-"`
	Code       int    `json:"code"`
	UserId     int64  `json:"userId"`
	CreateTime int64  `json:"createTime"`
	MusicA     string `json:"-"`
}
