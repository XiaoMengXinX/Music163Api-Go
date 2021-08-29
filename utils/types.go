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
