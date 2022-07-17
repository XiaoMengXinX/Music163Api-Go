package types

// DelEventData 删除动态 API 返回数据
type DelEventData struct {
	RawJson string `json:"-"`
	Message string `json:"message"`
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
}
