package types

// ShortURLData 返回数据
type ShortURLData struct {
	RawJson string `json:"-"`
	Code    int    `json:"code"`
	Data    struct {
		ShortUrl string `json:"shortUrl"`
	} `json:"data"`
	Message string `json:"message"`
}
