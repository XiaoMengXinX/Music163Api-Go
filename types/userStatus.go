package types

// UserStatusDetailData 用户状态详情数据
type UserStatusDetailData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Data    struct {
		Id       int         `json:"id"`
		UserId   int         `json:"userId"`
		Avatar   string      `json:"avatar"`
		UserName string      `json:"userName"`
		Song     interface{} `json:"song"`
		Content  struct {
			Type      string      `json:"type"`
			IconUrl   string      `json:"iconUrl"`
			Content   string      `json:"content"`
			ActionUrl interface{} `json:"actionUrl"`
		} `json:"content"`
		ExtInfo interface{} `json:"extInfo"`
	} `json:"data"`
	Message string `json:"message"`
}

// UserStatusSetData 用户状态设置数据
type UserStatusSetData struct {
	RawJson string
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}
