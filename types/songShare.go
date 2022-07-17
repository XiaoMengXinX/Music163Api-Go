package types

// SongShareData 分享歌曲数据
type SongShareData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Data    struct {
		IsWeChatMusicVideo   bool        `json:"isWeChatMusicVideo"`
		App                  interface{} `json:"app"`
		Business             interface{} `json:"business"`
		Scene                interface{} `json:"scene"`
		GroupId              interface{} `json:"groupId"`
		Type                 interface{} `json:"type"`
		Identify             interface{} `json:"identify"`
		TagBgColor           interface{} `json:"tagBgColor"`
		Tag                  interface{} `json:"tag"`
		SongAttribute        interface{} `json:"songAttribute"`
		CanShareWechatStatus bool        `json:"canShareWechatStatus"`
	} `json:"data"`
	Message string `json:"message"`
}
