package types

// SongsURLData 获取歌曲 URL API 的返回数据
type SongsURLData struct {
	RawJson string        `json:"-"`
	Data    []SongURLData `json:"data"`
	Code    int           `json:"code"`
}

// SongDownloadURLData SongDownloadURL API 的返回数据
type SongDownloadURLData struct {
	RawJson string      `json:"-"`
	Data    SongURLData `json:"data"`
	Code    int         `json:"code"`
}

// SongURLData 获取歌曲 URL API 的返回数据
type SongURLData struct {
	Id                 int         `json:"id"`
	Url                string      `json:"url"`
	Br                 int         `json:"br"`
	Size               int         `json:"size"`
	Md5                string      `json:"md5"`
	Code               int         `json:"code"`
	Expi               int         `json:"expi"`
	Type               string      `json:"type"`
	Gain               float64     `json:"gain"`
	Fee                int         `json:"fee"`
	Uf                 interface{} `json:"uf"`
	Paied              int         `json:"payed"`
	Flag               int         `json:"flag"`
	CanExtend          bool        `json:"canExtend"`
	FreeTrialInfo      interface{} `json:"freeTrialInfo"`
	Level              interface{} `json:"level"`
	EncodeType         interface{} `json:"encodeType"`
	FreeTrialPrivilege struct {
		ResConsumable  bool `json:"resConsumable"`
		UserConsumable bool `json:"userConsumable"`
	} `json:"freeTrialPrivilege"`
	FreeTimeTrialPrivilege struct {
		ResConsumable  bool `json:"resConsumable"`
		UserConsumable bool `json:"userConsumable"`
		Type           int  `json:"type"`
		RemainTime     int  `json:"remainTime"`
	} `json:"freeTimeTrialPrivilege"`
	UrlSource int `json:"urlSource"`
}
