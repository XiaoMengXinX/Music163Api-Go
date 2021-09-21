package utils

import "net/http"

// Headers 自定义 Headers 数据类型 (仅对于非 eapi 有效)
type Headers []struct {
	Name  string
	Value string
}

// RequestData 传入请求数据类型
type RequestData struct {
	Cookies []*http.Cookie
	Headers Headers
	Body    string
}

// EapiOption eapi 请求所需要的参数
type EapiOption struct {
	Json string
	Path string
	Url  string
}
