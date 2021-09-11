package types

// QrUnikeyData QrUnikey API 返回数据
type QrUnikeyData struct {
	RawJson string
	Code    int    `json:"code"`
	Unikey  string `json:"unikey"`
}
