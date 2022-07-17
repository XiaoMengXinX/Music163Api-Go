package types

// CloudBeanNumData 音乐人云豆数量
type CloudBeanNumData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ArtistId      int         `json:"artistId"`
		CloudBean     int         `json:"cloudBean"`
		MusicianLevel int         `json:"musicianLevel"`
		MaxCloudBean  int         `json:"maxCloudBean"`
		NeedPop       bool        `json:"needPop"`
		FlowStage     interface{} `json:"flowStage"`
		DynamicNews   interface{} `json:"dynamicNews"`
		Signed        bool        `json:"signed"`
	} `json:"data"`
}
