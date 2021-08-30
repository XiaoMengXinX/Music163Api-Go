package types

// UserInfoData 用户信息数据
type UserInfoData struct {
	Level     int `json:"level"`
	UserPoint struct {
		Balance      int   `json:"balance"`
		BlockBalance int   `json:"blockBalance"`
		Status       int   `json:"status"`
		UpdateTime   int64 `json:"updateTime"`
		UserId       int   `json:"userId"`
		Version      int   `json:"version"`
	} `json:"userPoint"`
	MobileSign       bool   `json:"mobileSign"`
	PcSign           bool   `json:"pcSign"`
	Viptype          int    `json:"viptype"`
	Expiretime       int64  `json:"expiretime"`
	BackupExpireTime int64  `json:"backupExpireTime"`
	Storeurl         string `json:"storeurl"`
	MallDesc         string `json:"mallDesc"`
	StoreTitle       string `json:"storeTitle"`
	Pubwords         string `json:"pubwords"`
	GameConfig       struct {
		ExpirationTime       int64  `json:"expirationTime"`
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
