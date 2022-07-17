package types

// UserInfoData 用户信息数据
type UserInfoData struct {
	RawJson   string `json:"-"`
	Level     int    `json:"level"`
	UserPoint struct {
		Balance      int `json:"balance"`
		BlockBalance int `json:"blockBalance"`
		Status       int `json:"status"`
		UpdateTime   int `json:"updateTime"`
		UserId       int `json:"userId"`
		Version      int `json:"version"`
	} `json:"userPoint"`
	MobileSign       bool   `json:"mobileSign"`
	PcSign           bool   `json:"pcSign"`
	Viptype          int    `json:"viptype"`
	Expiretime       int    `json:"expiretime"`
	BackupExpireTime int    `json:"backupExpireTime"`
	Storeurl         string `json:"storeurl"`
	MallDesc         string `json:"mallDesc"`
	StoreTitle       string `json:"storeTitle"`
	Pubwords         string `json:"pubwords"`
	GameConfig       struct {
		ExpirationTime       int    `json:"expirationTime"`
		GameCenterEntryTitle string `json:"gameCenterEntryTitle"`
		GameCenterPic        string `json:"gameCenterPic"`
		GameCenterReddot     int    `json:"gameCenterReddot"`
		GameCenterText       string `json:"gameCenterText"`
		GameCenterUrl        string `json:"gameCenterUrl"`
	} `json:"gameConfig"`
	RingConfig interface{} `json:"ringConfig"`
	FmConfig   struct {
		Image string `json:"image"`
		Text  string `json:"text"`
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"fmConfig"`
	TicketConfig interface{} `json:"ticketConfig"`
	Code         int         `json:"code"`
}
