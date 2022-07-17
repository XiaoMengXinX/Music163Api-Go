package types

// QrUnikeyData QrUnikey API 返回数据
type QrUnikeyData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Unikey  string `json:"unikey"`
}
