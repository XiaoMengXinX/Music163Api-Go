package utils

// Cookies 用户 cookies 数据类型
type Cookies []struct {
	Key   string
	Value string
}

// RequestData 传入请求数据类型
type RequestData struct {
	Cookies Cookies
	Body    string
}

// EapiOption eapi 请求所需要的参数
type EapiOption struct {
	Json string
	Path string
	Url  string
}

// Headers 返回内容的 Headers
type Headers struct {
	CacheControl  []string `json:"Cache-Control"`
	ContentType   []string `json:"Content-Type"`
	Date          []string `json:"Date"`
	Expires       []string `json:"Expires"`
	GwThread      []string `json:"Gw-Thread"`
	GwTime        []string `json:"Gw-Time"`
	MconfigBucket []string `json:"Mconfig-Bucket"`
	Server        []string `json:"Server"`
	SetCookie     []string `json:"Set-Cookie"`
	XFromSrc      []string `json:"X-From-Src"`
	XTraceid      []string `json:"X-Traceid"`
	XVia          []string `json:"X-Via"`
}
