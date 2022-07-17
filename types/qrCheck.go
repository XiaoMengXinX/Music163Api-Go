package types

// QrCheckData QrCheck API 返回数据
type QrCheckData struct {
	RawJson   string `json:"-"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatarUrl"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}
